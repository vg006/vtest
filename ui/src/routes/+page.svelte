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

<div>
    <div>
        <h1>Vtest</h1>
    </div>
    <div>
        <form>
            <input
                name="port"
                type="number"
                placeholder="PORT"
                bind:value={port}
            />
            <input
                name="endpoint"
                type="text"
                placeholder="ENDPOINT"
                bind:value={endpoint}
            />
            <select name="method" bind:value={method}>
                <option value="ALL" selected>ALL</option>
                <option value="GET">GET</option>
                <option value="POST">POST</option>
                <option value="PUT">PUT</option>
                <option value="DELETE">DELETE</option>
            </select>
            <button onclick={handleSubmit}>Analyze</button>
        </form>
    </div>
    <div>
        {#if app.loading}
            <div>
                <p>Loading...</p>
            </div>
        {:else if app.errorMsg}
            <div>
                <p>Error: {app.errorMsg}</p>
            </div>
        {:else if app.resData}
            <div>
                <div>
                    <h2>Response</h2>
                    <p>Status: {app.resData.StatusCode}</p>
                </div>
                {#each app.resData.Headers as header}
                    <div>
                        <h2>Headers</h2>
                        <ul>
                            <li>{header.Key}: {header.Value}</li>
                        </ul>
                    </div>
                {/each}
            </div>
        {/if}
    </div>
</div>
