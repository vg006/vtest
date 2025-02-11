<script lang="ts">
    import type { Response, Header } from "$lib/index";

    let port = $state("");
    let endpoint = $state("");
    let method = $state("GET");

    let app: {
        loading: boolean;
        errorMsg: string;
        resData: Response | null;
        securityScore?: number;
    } = $state({
        loading: false,
        errorMsg: "",
        resData: null,
        securityScore: undefined,
    });

    let routes: {
        port: string;
        endpoint: string;
        method: string;
    }[] = $state([]);

    const addRoute = (port: string, endpoint: string, method: string) => {
        const newRoute = { port, endpoint, method };
        const exists = routes.some(
            (route) =>
                route.port === port &&
                route.endpoint === endpoint &&
                route.method === method,
        );
        if (!exists) {
            routes = [...routes, newRoute];
        }
    };

    const generateScore = (headers: Header[]) => {
        const weightedScores: {
            [key: string]: number;
        } = {
            "Strict-Transport-Security": 0.2,
            "Content-Security-Policy": 0.2,
            "X-Frame-Options": 0.15,
            "X-Content-Type-Options": 0.15,
            "X-XSS-Protection": 0.15,
            "Referrer-Policy": 0.15,
            "Permission-Policy": 0.1,
        };

        let totalScore = 0;
        headers.forEach((header) => {
            if (header.Presence && weightedScores[header.Key]) {
                totalScore += weightedScores[header.Key] * 100;
            }
        });

        return Math.round(totalScore);
    };

    const fetchDetails = async () => {
        if (!/^\d+$/.test(port) || port.length === 0) {
            app.errorMsg = "Port must be numeric.";
            console.log("Port must be numeric.");
            return;
        }
        app.errorMsg = "";

        app.loading = true;
        const payload = {
            port: String(port),
            endpoint: "/" + endpoint,
            method,
        };
        try {
            const res = await fetch(`http://localhost:8080/api/details`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload),
            });
            if (!res.ok) {
                throw new Error(`Server responded with ${res.status}`);
            }
            addRoute(port, endpoint, method);
            app.loading = false;
            app.resData = await res.json();
            if (app.resData !== null) {
                app.securityScore = generateScore(app.resData.Headers);
                addRoute(port, endpoint, method);
            } else {
                app.securityScore = undefined;
                app.errorMsg = "Failed to fetch data";
            }
        } catch (err: any) {
            app.loading = false;
            app.errorMsg = err.message;
        }
    };
</script>

<div
    class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 py-12 px-4 transition-all duration-300"
>
    <div class="max-w-6xl mx-auto">
        <div class="flex flex-col">
            <div class="animate-fade-in mb-6 text-left">
                <h1
                    class="text-6xl font-bold bg-gradient-to-b from-orange-600 to-pink-400 bg-clip-text text-transparent hover:scale-105 transition-transform duration-300"
                >
                    Vtest
                </h1>
            </div>
            <form class="flex flex-wrap items-start gap-4 mb-12">
                <div class="flex items-center group">
                    <div
                        class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-3 group-hover:border-orange-500 transition-all duration-300"
                    >
                        <span
                            class="text-gray-600 font-semibold group-hover:text-orange-500 group-focus:text-orange-600"
                            >:</span
                        >
                    </div>
                    <input
                        name="port"
                        type="text"
                        inputmode="numeric"
                        pattern="[0-9]*"
                        placeholder="PORT"
                        bind:value={port}
                        class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-3 w-40 focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                    />
                </div>
                <div class="flex items-center group">
                    <div
                        class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-3 group-hover:border-orange-500 transition-all duration-300"
                    >
                        <span
                            class="text-gray-600 font-semibold group-hover:text-orange-500 group-focus:text-orange-600"
                            >/</span
                        >
                    </div>
                    <input
                        name="endpoint"
                        type="text"
                        placeholder="ENDPOINT"
                        bind:value={endpoint}
                        class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-3 w-80 focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300"
                    />
                </div>
                <select
                    name="method"
                    placeholder="METHOD"
                    bind:value={method}
                    class="border-2 border-gray-200 rounded-lg px-6 py-3 focus:ring-2 focus:ring-orange-500 focus:border-transparent outline-none transition-all duration-300"
                >
                    <option value="GET">GET</option>
                    <option value="POST">POST</option>
                    <option value="PUT">PUT</option>
                    <option value="DELETE">DELETE</option>
                </select>
                <button
                    onclick={fetchDetails}
                    class="bg-gradient-to-r from-orange-500 to-pink-500 text-white px-8 py-3 rounded-lg hover:from-orange-600 hover:to-pink-600 transform hover:-translate-y-0.5 transition-all duration-300 shadow-md hover:shadow-lg font-medium"
                >
                    Analyze
                </button>
            </form>

            <div class="flex gap-6">
                <div class="w-1/3">
                    <div class="bg-white p-6 rounded-xl shadow-lg mb-6">
                        <h2 class="text-xl font-semibold mb-4">
                            Recent Routes
                        </h2>
                        <div class="space-y-2">
                            {#each routes as route}
                                <div class="p-3 bg-gray-50 rounded">
                                    <span class="font-medium"
                                        >{route.method}</span
                                    >
                                    <span class="text-gray-600"
                                        >:{route.port}/{route.endpoint}</span
                                    >
                                </div>
                            {/each}
                        </div>
                    </div>
                </div>

                <div class="w-2/3">
                    {#if app.loading}
                        <div class="flex justify-center p-12">
                            <div
                                class="animate-spin rounded-full h-16 w-16 border-4 border-orange-500 border-t-transparent"
                            ></div>
                        </div>
                    {:else if app.errorMsg}
                        <div
                            class="bg-red-50 text-red-700 p-8 rounded-lg shadow-md animate-fade-in"
                        >
                            <p class="flex items-center">
                                <svg
                                    class="w-6 h-6 mr-3"
                                    fill="currentColor"
                                    viewBox="0 0 20 20"
                                >
                                    <path
                                        fill-rule="evenodd"
                                        d="M10 18a8 8 0 100-16 8 8 0 000 16zM8.707 7.293a1 1 0 00-1.414 1.414L8.586 10l-1.293 1.293a1 1 0 101.414 1.414L10 11.414l1.293 1.293a1 1 0 001.414-1.414L11.414 10l1.293-1.293a1 1 0 00-1.414-1.414L10 8.586 8.707 7.293z"
                                        clip-rule="evenodd"
                                    />
                                </svg>
                                Error: {app.errorMsg}
                            </p>
                        </div>
                    {:else if app.resData}
                        <div class="space-y-6">
                            <div
                                class="bg-white p-10 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300"
                            >
                                <h2
                                    class="text-3xl font-semibold mb-8 text-gray-800"
                                >
                                    Security Metrics
                                </h2>
                                <div class="mb-6">
                                    <p
                                        class="text-gray-600 text-sm uppercase tracking-wider mb-2"
                                    >
                                        Status Code
                                    </p>
                                    <p
                                        class="text-4xl font-bold {app.resData
                                            .StatusCode >= 500
                                            ? 'text-red-600'
                                            : app.resData.StatusCode >= 400
                                              ? 'text-orange-500'
                                              : app.resData.StatusCode >= 300
                                                ? 'text-blue-500'
                                                : app.resData.StatusCode >= 200
                                                  ? 'text-green-500'
                                                  : 'text-gray-600'}"
                                    >
                                        {app.resData.StatusCode}
                                    </p>
                                </div>
                                {#if app.securityScore !== undefined}
                                    <div class="mb-6">
                                        <p
                                            class="text-gray-600 text-sm uppercase tracking-wider mb-2"
                                        >
                                            Security Score
                                        </p>
                                        <p
                                            class="text-4xl font-bold {app.securityScore >=
                                            80
                                                ? 'text-green-500'
                                                : app.securityScore >= 60
                                                  ? 'text-yellow-500'
                                                  : app.securityScore >= 40
                                                    ? 'text-orange-500'
                                                    : 'text-red-500'}"
                                        >
                                            {app.securityScore}%
                                        </p>
                                    </div>
                                {/if}
                            </div>

                            <div
                                class="bg-white p-10 rounded-xl shadow-lg hover:shadow-xl transition-all duration-300"
                            >
                                <h2
                                    class="text-3xl font-semibold mb-8 text-gray-800"
                                >
                                    Headers
                                </h2>
                                <div class="divide-y divide-gray-100">
                                    {#each app.resData.Headers as header}
                                        <details class="py-4 group">
                                            <summary
                                                class="flex justify-between items-center cursor-pointer hover:bg-gray-50 p-3 rounded-lg transition-colors duration-300"
                                            >
                                                <span
                                                    class="font-medium text-gray-800 group-hover:text-orange-500 transition-colors duration-300"
                                                    >{header.Key}</span
                                                >
                                                <span
                                                    class="text-gray-600 group-hover:text-pink-500 transition-colors duration-300"
                                                    >{header.Value}</span
                                                >
                                            </summary>
                                            <div
                                                class="pt-4 pl-8 text-gray-600 bg-gray-50 mt-3 rounded-lg p-4 transition-all duration-300 origin-top transform scale-y-0 opacity-0 group-open:scale-y-100 group-open:opacity-100 flex flex-row justify-between [&[open]>.group-open\:scale-y-100]:scale-y-100 [&[open]>.group-open\:opacity-100]:opacity-100"
                                            >
                                                <p class="text-sm">
                                                    {header.Usage}
                                                </p>
                                                <div
                                                    class="flex items-center gap-2"
                                                >
                                                    <span
                                                        class="h-3 w-3 rounded-full {header.Level ===
                                                        'Critical'
                                                            ? 'bg-red-500'
                                                            : header.Level ===
                                                                'High'
                                                              ? 'bg-orange-500'
                                                              : header.Level ===
                                                                  'Medium'
                                                                ? 'bg-yellow-500'
                                                                : 'bg-blue-500'}"
                                                    ></span>
                                                    <p>{header.Level}</p>
                                                </div>
                                            </div>
                                        </details>
                                    {/each}
                                </div>
                            </div>
                        </div>
                    {/if}
                </div>
            </div>
        </div>
    </div>
</div>
