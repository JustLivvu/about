<script lang="ts">
    import Icon from "@iconify/svelte";
</script>

<div class="followers-container">
    <div class="empty-state">
        <p class="title">No followers yet</p>
        <p class="subtitle">When someone follows this profile, they will appear here.</p>
    </div>
</div>

<style>
    .followers-container {
        padding: 24px 0;
        width: 100%;
    }

    .empty-state {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
        color: #8588a3;
    }

    .empty-state :global(svg) {
        width: 48px;
        height: 48px;
        margin-bottom: 12px;
        color: #585b70;
    }

    .title {
        font-size: 1rem;
        font-weight: 600;
        color: #c9d2f0;
        margin: 0 0 6px 0;
    }

    .subtitle {
        font-size: 0.9rem;
        color: #8588a3;
        margin: 0;
    }
</style>
