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

    let recentRoutes: {
        port: string;
        endpoint: string;
        method: string;
    }[] = $state([]);

    let showAddCollectionModal = $state(false);
    let newCollectionName = $state("");
    let newCollectionPort = $state();
    let newCollectionEndpoint = $state("");

    let collections: {
        name: string;
        routes: {
            port: number;
            endpoint: string;
            method: string;
        }[];
    }[] = $state([]);

    const addCollection = async () => {
        if (newCollectionName.trim() !== "") {
            if (newCollectionEndpoint === "") {
                newCollectionEndpoint = "vtest/routes";
            }
            collections = [
                ...collections,
                {
                    name: newCollectionName,
                    routes: await fetchRoutes(),
                },
            ];
            newCollectionName = "";
            newCollectionPort = 0;
            newCollectionEndpoint = "";
            showAddCollectionModal = false;
        }
    };

    const fetchRoutes = async () => {
        try {
            const res = await fetch(`http://localhost:8080/api/routes`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify({
                    port: newCollectionPort,
                    endpoint: "/" + newCollectionEndpoint,
                }),
            });
            if (!res.ok) {
                throw new Error(`Server responded with ${res.status}`);
            }
            let resData = await res.json();
            resData.forEach(
                (route: { port: number; endpoint: string; method: string }) => {
                    route.port = newCollectionPort as number;
                },
            );
            return resData;
        } catch (err: any) {
            console.log(err.message);
            return [];
        }
    };

    const addRoute = (port: string, endpoint: string, method: string) => {
        const newRoute = { port, endpoint, method };
        const exists = recentRoutes.some(
            (route) =>
                route.port === port &&
                route.endpoint === endpoint &&
                route.method === method,
        );
        if (!exists) {
            recentRoutes = [...recentRoutes, newRoute];
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
        if (endpoint[0] === "/") {
            endpoint = endpoint.slice(1);
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

<div class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 p-8">
    <div class="max-w-6xl mx-auto">
        <div class="flex flex-col space-y-8">
            <div class="animate-fade-in">
                <h1
                    class="text-5xl font-bold bg-gradient-to-r from-orange-600 to-pink-400 bg-clip-text text-transparent"
                >
                    Vtest
                </h1>
                <p class="text-gray-600">
                    A tool to analyze security metrics of your APIs.
                </p>
            </div>

            <form class="flex flex-wrap items-start gap-3">
                <div class="flex items-center group">
                    <div
                        class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-2.5 group-hover:border-orange-500 transition-all duration-300"
                    >
                        <span
                            class="text-gray-600 font-semibold group-hover:text-orange-500"
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
                        class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-2.5 w-32 focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300 [appearance:textfield] [&::-webkit-outer-spin-button]:appearance-none [&::-webkit-inner-spin-button]:appearance-none"
                    />
                </div>

                <div class="flex items-center group">
                    <div
                        class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-2.5 group-hover:border-orange-500 transition-all duration-300"
                    >
                        <span
                            class="text-gray-600 font-semibold group-hover:text-orange-500"
                            >/</span
                        >
                    </div>
                    <input
                        name="endpoint"
                        type="text"
                        placeholder="ENDPOINT"
                        bind:value={endpoint}
                        class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-2.5 w-64 focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300"
                    />
                </div>

                <select
                    name="method"
                    bind:value={method}
                    class="border-2 border-gray-200 rounded-lg px-4 py-2.5 focus:ring-2 focus:ring-orange-500 focus:border-transparent outline-none transition-all duration-300"
                >
                    <option value="GET">GET</option>
                    <option value="POST">POST</option>
                    <option value="PUT">PUT</option>
                    <option value="DELETE">DELETE</option>
                </select>

                <button
                    onclick={fetchDetails}
                    class="bg-gradient-to-r from-orange-500 to-pink-500 text-white px-6 py-2.5 rounded-lg hover:from-orange-600 hover:to-pink-600 transform hover:-translate-y-0.5 transition-all duration-300 shadow-md hover:shadow-lg font-medium"
                >
                    Analyze
                </button>
            </form>

            <div class="flex gap-8">
                <div class="w-1/3 space-y-6">
                    <div
                        class="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-all duration-300"
                    >
                        <h2 class="text-xl font-semibold mb-4">
                            Recent Routes
                        </h2>
                        <div class="space-y-2">
                            {#each recentRoutes as route}
                                <div
                                    class="flex justify-between items-center p-3 bg-gray-50 rounded hover:bg-gray-100 transition-all duration-300"
                                >
                                    <div>
                                        <span class="font-medium"
                                            >{route.method}</span
                                        >
                                        <span class="text-gray-600"
                                            >:{route.port}/{route.endpoint}</span
                                        >
                                    </div>
                                    <button
                                        aria-label="Analyze route"
                                        class="bg-orange-500 text-white p-1.5 rounded-full hover:bg-orange-600 transition-colors duration-300"
                                        onclick={() => {
                                            port = route.port.toString();
                                            endpoint = route.endpoint;
                                            method = route.method;
                                            fetchDetails();
                                        }}
                                    >
                                        <svg
                                            class="w-4 h-4"
                                            fill="none"
                                            stroke="currentColor"
                                            viewBox="0 0 24 24"
                                        >
                                            <path
                                                stroke-linecap="round"
                                                stroke-linejoin="round"
                                                stroke-width="2"
                                                d="M9 5l7 7-7 7"
                                            />
                                        </svg>
                                    </button>
                                </div>
                            {/each}
                        </div>
                    </div>

                    <div
                        class="bg-white p-6 rounded-xl shadow-md hover:shadow-lg transition-all duration-300"
                    >
                        <div class="flex justify-between items-center mb-4">
                            <h2 class="text-xl font-semibold">
                                My Collections
                            </h2>
                            <button
                                aria-label="Add collection"
                                class="bg-orange-500 text-white p-1.5 rounded-full hover:bg-orange-600 transition-colors duration-300"
                                onclick={() => (showAddCollectionModal = true)}
                            >
                                <svg
                                    class="w-5 h-5"
                                    fill="none"
                                    stroke="currentColor"
                                    viewBox="0 0 24 24"
                                >
                                    <path
                                        stroke-width="2"
                                        d="M12 6v6m0 0v6m0-6h6m-6 0H6"
                                    />
                                </svg>
                            </button>
                        </div>

                        <div class="space-y-3">
                            {#each collections as collection, i}
                                <details
                                    class="bg-gray-50 rounded-lg transition-all duration-300"
                                >
                                    <summary
                                        class="p-3 font-medium cursor-pointer flex justify-between items-center"
                                    >
                                        <span>{collection.name}</span>
                                        <button
                                            aria-label="Delete collection"
                                            class="text-red-500 hover:text-red-700 transition-colors duration-300"
                                            onclick={() =>
                                                (collections =
                                                    collections.filter(
                                                        (_, index) =>
                                                            index !== i,
                                                    ))}
                                        >
                                            <svg
                                                class="w-5 h-5"
                                                fill="none"
                                                stroke="currentColor"
                                                viewBox="0 0 24 24"
                                            >
                                                <path
                                                    stroke-linecap="round"
                                                    stroke-linejoin="round"
                                                    stroke-width="2"
                                                    d="M19 7l-.867 12.142A2 2 0 0116.138 21H7.862a2 2 0 01-1.995-1.858L5 7m5 4v6m4-6v6m1-10V4a1 1 0 00-1-1h-4a1 1 0 00-1 1v3M4 7h16"
                                                />
                                            </svg>
                                        </button>
                                    </summary>
                                    <div class="p-3 space-y-2">
                                        {#each collection.routes as route}
                                            <div
                                                class="flex justify-between items-center p-2 bg-white rounded-lg"
                                            >
                                                <div>
                                                    <span class="font-medium"
                                                        >{route.method}</span
                                                    >
                                                    <span class="text-gray-600"
                                                        >:{route.port}{route.endpoint}</span
                                                    >
                                                </div>
                                                <button
                                                    aria-label="Analyze route"
                                                    class="bg-orange-500 text-white p-1.5 rounded-full hover:bg-orange-600 transition-colors duration-300"
                                                    onclick={() => {
                                                        port =
                                                            route.port.toString();
                                                        endpoint =
                                                            route.endpoint;
                                                        method = route.method;
                                                        fetchDetails();
                                                    }}
                                                >
                                                    <svg
                                                        class="w-4 h-4"
                                                        fill="none"
                                                        stroke="currentColor"
                                                        viewBox="0 0 24 24"
                                                    >
                                                        <path
                                                            stroke-linecap="round"
                                                            stroke-linejoin="round"
                                                            stroke-width="2"
                                                            d="M9 5l7 7-7 7"
                                                        />
                                                    </svg>
                                                </button>
                                            </div>
                                        {/each}
                                    </div>
                                </details>
                            {/each}
                        </div>
                    </div>

                    {#if showAddCollectionModal}
                        <div
                            class="fixed inset-0 bg-black bg-opacity-50 flex items-center justify-center z-50"
                        >
                            <div class="bg-white p-6 rounded-xl w-96 shadow-xl">
                                <h3 class="text-xl font-semibold mb-4">
                                    Add New Collection
                                </h3>
                                <div class="space-y-4">
                                    <input
                                        type="text"
                                        placeholder="Collection Name"
                                        bind:value={newCollectionName}
                                        class="w-full border-2 border-gray-200 rounded-lg p-2.5 focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent"
                                    />
                                    <div class="flex items-center group">
                                        <div
                                            class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-2.5 group-hover:border-orange-500 transition-all duration-300"
                                        >
                                            <span
                                                class="text-gray-600 font-semibold group-hover:text-orange-500"
                                                >:</span
                                            >
                                        </div>
                                        <input
                                            name="port"
                                            type="text"
                                            inputmode="numeric"
                                            pattern="[0-9]*"
                                            placeholder="PORT"
                                            bind:value={newCollectionPort}
                                            class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-2.5 w-full focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300"
                                        />
                                    </div>
                                    <div class="flex items-center group">
                                        <div
                                            class="flex rounded-l-lg bg-gray-200 border-2 border-r-0 border-gray-200 px-3 py-2.5 group-hover:border-orange-500 transition-all duration-300"
                                        >
                                            <span
                                                class="text-gray-600 font-semibold group-hover:text-orange-500"
                                                >/</span
                                            >
                                        </div>
                                        <input
                                            name="endpoint"
                                            type="text"
                                            placeholder="ENDPOINT"
                                            bind:value={newCollectionEndpoint}
                                            class="text-left rounded-r-lg border-2 border-gray-200 px-4 py-2.5 w-full focus:outline-none focus:ring focus:ring-orange-500 focus:border-transparent group-hover:border-2 group-hover:border-orange-500 transition-all duration-300"
                                        />
                                    </div>
                                </div>

                                <div class="flex justify-end gap-3 mt-6">
                                    <button
                                        class="px-4 py-2 bg-gray-200 rounded-lg hover:bg-gray-300 transition-colors duration-300"
                                        onclick={() =>
                                            (showAddCollectionModal = false)}
                                    >
                                        Cancel
                                    </button>
                                    <button
                                        class="px-4 py-2 bg-orange-500 text-white rounded-lg hover:bg-orange-600 transition-colors duration-300"
                                        onclick={addCollection}
                                    >
                                        Add
                                    </button>
                                </div>
                            </div>
                        </div>
                    {/if}
                </div>

                <div class="w-2/3">
                    {#if app.loading}
                        <div class="flex justify-center items-center h-64">
                            <div
                                class="animate-spin rounded-full h-16 w-16 border-4 border-orange-500 border-t-transparent"
                            ></div>
                        </div>
                    {:else if app.errorMsg}
                        <div
                            class="bg-red-50 text-red-700 p-6 rounded-xl shadow-md animate-fade-in"
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
                                class="bg-white p-8 rounded-xl shadow-md hover:shadow-lg transition-all duration-300"
                            >
                                <h2
                                    class="text-2xl font-semibold mb-6 text-gray-800"
                                >
                                    Security Metrics
                                </h2>
                                <div class="flex gap-12">
                                    <div>
                                        <p
                                            class="text-gray-600 text-sm uppercase tracking-wider mb-2"
                                        >
                                            Status Code
                                        </p>
                                        <p
                                            class="text-4xl font-bold {app
                                                .resData.StatusCode >= 500
                                                ? 'text-red-600'
                                                : app.resData.StatusCode >= 400
                                                  ? 'text-orange-500'
                                                  : app.resData.StatusCode >=
                                                      300
                                                    ? 'text-blue-500'
                                                    : app.resData.StatusCode >=
                                                        200
                                                      ? 'text-green-500'
                                                      : 'text-gray-600'}"
                                        >
                                            {app.resData.StatusCode}
                                        </p>
                                    </div>
                                    {#if app.securityScore !== undefined}
                                        <div>
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
                            </div>

                            <div
                                class="bg-white p-8 rounded-xl shadow-md hover:shadow-lg transition-all duration-300"
                            >
                                <h2
                                    class="text-2xl font-semibold mb-6 text-gray-800"
                                >
                                    Headers
                                </h2>
                                <div class="space-y-2">
                                    {#each app.resData.Headers as header}
                                        <details
                                            class="group bg-gray-50 rounded-lg hover:bg-gray-100 transition-all duration-300"
                                        >
                                            <summary
                                                class="flex justify-between items-center p-4 cursor-pointer"
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
                                                class="p-4 bg-white rounded-b-lg"
                                            >
                                                <div
                                                    class="flex justify-between items-center"
                                                >
                                                    <p
                                                        class="text-sm text-gray-600"
                                                    >
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
                                                        <p
                                                            class="text-sm font-medium"
                                                        >
                                                            {header.Level}
                                                        </p>
                                                    </div>
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
