<script lang="ts">
    let searchQuery = $state('');
</script>

<div class="followers-container">
    <input
        type="text"
        class="search-input"
        placeholder="Search for followers"
        bind:value={searchQuery}
    />

    <div class="empty-state">
        <p class="title">No followers yet</p>
        <p class="subtitle">When someone follows this profile, they will appear here.</p>
    </div>
</div>

<style>
    .followers-container {
        padding: 24px 0;
        width: 100%;
        display: flex;
        flex-direction: column;
        gap: 24px;
    }

    .search-input {
        width: 100%;
        box-sizing: border-box;
        padding: 10px 14px;
        background-color: #181825;
        border: 1px solid #3f4153;
        border-radius: 4px;
        color: #c9d2f0;
        font-size: 0.9rem;
        font-family: inherit;
        outline: none;
        transition: border-color 0.15s ease;
    }

    .search-input:focus {
        border-color: #7aa2f7;
    }

    .search-input::placeholder {
        color: #8588a3;
        font-weight: 300;
    }

    .empty-state {
        display: flex;
        flex-direction: column;
        align-items: center;
        justify-content: center;
        text-align: center;
        color: #8588a3;
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
