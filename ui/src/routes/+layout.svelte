<script lang="ts">
    import "../app.css";
    let { children } = $props();
</script>

<div class="custom-container px-40 tab:px-40">
    {@render children()}
</div>
