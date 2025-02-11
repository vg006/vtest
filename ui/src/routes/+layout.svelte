<script lang="ts">
    import "../app.css";
    let { children } = $props();
</script>

<div class="custom-container flex flex-row flex-nowrap">
    <div
        class="w-1/20 bg-gray-400 transition-transform duration-300 ease-in-out transform origin-left hover:scale-x-150"
    ></div>
    <div class="w-full">
        {@render children()}
    </div>
</div>
