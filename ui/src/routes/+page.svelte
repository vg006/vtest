<script lang="ts">
    import type { Response } from "$lib/index";

    let port = $state("");
    let endpoint = $state("");
    let method = $state("ALL");

    let app: {
        loading: boolean;
        errorMsg: string;
        resData: Response | null;
    } = $state({
        loading: false,
        errorMsg: "",
        resData: null,
    });

    const handleSubmit = async () => {
        if (!/^\d+$/.test(port) || port.length === 0) {
            app.errorMsg = "Port must be numeric.";
            console.log("Port must be numeric.");
            return;
        }
        if (!/^\/.*/.test(endpoint) || endpoint.length === 0) {
            app.errorMsg = "Endpoint must start with '/'.";
            console.log("Endpoint must start with '/'.");
            return;
        }
        app.errorMsg = "";

        app.loading = true;
        const payload = { port: String(port), endpoint, method };
        try {
            const res = await fetch(`http://localhost:8080/api/details`, {
                method: "POST",
                headers: { "Content-Type": "application/json" },
                body: JSON.stringify(payload),
            });
            if (!res.ok) {
                throw new Error(`Server responded with ${res.status}`);
            }
            app.loading = false;
            app.resData = await res.json();
        } catch (err: any) {
            app.loading = false;
            app.errorMsg = err.message;
        }
    };
</script>

<div
    class="min-h-screen bg-gradient-to-br from-gray-50 to-gray-100 py-8 transition-all duration-300"
>
    <div class="text-center mb-12 animate-fade-in">
        <h1
            class="text-5xl font-bold bg-gradient-to-r from-orange-400 via-pink-500 to-purple-600 bg-clip-text text-transparent hover:scale-105 transition-transform duration-300"
        >
            Vtest
        </h1>
    </div>
    <div class="max-w-4xl mx-auto px-4">
        <form
            class="flex items-center gap-4 mb-8 bg-white p-6 rounded-lg shadow-lg hover:shadow-xl transition-all duration-300"
        >
            <div class="flex items-center group">
                <span
                    class="mr-1 text-gray-600 group-hover:text-orange-500 transition-colors duration-300"
                    >:</span
                >
                <input
                    name="port"
                    type="number"
                    placeholder="PORT"
                    bind:value={port}
                    class="border border-gray-200 rounded-lg px-4 py-2 focus:ring-2 focus:ring-orange-500 focus:border-transparent outline-none transition-all duration-300"
                />
            </div>
            <div class="flex items-center group">
                <span
                    class="mr-1 text-gray-600 group-hover:text-orange-500 transition-colors duration-300"
                    >/</span
                >
                <input
                    name="endpoint"
                    type="text"
                    placeholder="ENDPOINT"
                    bind:value={endpoint}
                    class="border border-gray-200 rounded-lg px-4 py-2 focus:ring-2 focus:ring-orange-500 focus:border-transparent outline-none transition-all duration-300"
                />
            </div>
            <select
                name="method"
                bind:value={method}
                class="border border-gray-200 rounded-lg px-4 py-2 focus:ring-2 focus:ring-orange-500 focus:border-transparent outline-none transition-all duration-300"
            >
                <option value="ALL" selected>ALL</option>
                <option value="GET">GET</option>
                <option value="POST">POST</option>
                <option value="PUT">PUT</option>
                <option value="DELETE">DELETE</option>
            </select>
            <button
                onclick={handleSubmit}
                class="bg-gradient-to-r from-orange-500 to-pink-500 text-white px-6 py-2 rounded-lg hover:from-orange-600 hover:to-pink-600 transform hover:-translate-y-0.5 transition-all duration-300 shadow-md hover:shadow-lg"
            >
                Analyze
            </button>
        </form>
        <div>
            {#if app.loading}
                <div class="flex justify-center p-8">
                    <div
                        class="animate-spin rounded-full h-12 w-12 border-4 border-orange-500 border-t-transparent"
                    ></div>
                </div>
            {:else if app.errorMsg}
                <div
                    class="bg-red-50 text-red-700 p-6 rounded-lg shadow-md animate-fade-in"
                >
                    <p class="flex items-center">
                        <svg
                            class="w-5 h-5 mr-2"
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
                <div class="grid grid-cols-2 gap-8 animate-fade-in">
                    <div
                        class="bg-white p-8 rounded-lg shadow-lg hover:shadow-xl transition-all duration-300"
                    >
                        <h2 class="text-2xl font-semibold mb-6 text-gray-800">
                            Security Metrics
                        </h2>
                        <div class="mb-4">
                            <p
                                class="text-gray-600 text-sm uppercase tracking-wide"
                            >
                                Status Code
                            </p>
                            <p
                                class="text-3xl font-bold bg-gradient-to-r from-orange-400 to-pink-500 bg-clip-text text-transparent"
                            >
                                {app.resData.StatusCode}
                            </p>
                        </div>
                    </div>
                    <div
                        class="bg-white p-8 rounded-lg shadow-lg hover:shadow-xl transition-all duration-300"
                    >
                        <h2 class="text-2xl font-semibold mb-6 text-gray-800">
                            Headers
                        </h2>
                        <div class="divide-y divide-gray-100">
                            {#each app.resData.Headers as header}
                                <details class="py-3 group">
                                    <summary
                                        class="flex justify-between items-center cursor-pointer hover:bg-gray-50 p-2 rounded-lg transition-colors duration-300"
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
                                        class="pt-3 pl-6 text-gray-600 bg-gray-50 mt-2 rounded-lg"
                                    >
                                        <p class="text-sm">
                                            Header usage information goes here
                                        </p>
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
