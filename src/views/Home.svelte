<script lang="ts">
    import { api } from '../lib/api';
    import Icon from "@iconify/svelte"

    async function getAvatar() {
        const response = await api('/media/avatar.jpg');

        if (!response.ok) {
            throw new Error('No avatar for ya');
        }

        const blob = await response.blob();
        return URL.createObjectURL(blob);
    }

    let avatarUrl = '';

    getAvatar().then((url) => {
        avatarUrl = url;
    });
</script>

<div class="home_main">
    <div class="info_section">
        <div class="profile_picture">
            <img src={avatarUrl} alt="Avatar" />
        </div>

        <div class="basic_info">
            <span class="username">Livvya</span>
            <span class="pronouns">livvya · she/her</span>
        </div>

        <div class="follow_stats">
            <span>0 followers, 0 following</span>

            <button class="more-btn">
                <Icon icon="mdi:dots-horizontal" />
            </button>
        </div>

        <div class="separator"></div>

        <div class="links">
            <div class="link">
                <Icon icon="mdi:link-variant" />
                <a
                    href="https://github.com/justlivvu"
                    target="_blank"
                    rel="noopener noreferrer"
                >
                    github.com/justlivvu
                </a>
            </div>
        </div>
        <div class="separator"></div>
        <div class="description">
            <span>Hey! I'm Livvya, Developer and maintainer of Picobio, Feel free to fork my project or dm me on discord or here. Dont forget to star Picobio on GitHub!</span>
        </div>
        <div class="separator"></div>
        <div class="profile_creation">
            <Icon icon="material-symbols-light:calendar-today"/>
            <span>Joined on Jun 14, 2026</span>
        </div>
    </div>

    <div class="actual_content">
        <div class="navbar">
            <span class="tab-chosen">
                <Icon icon="material-symbols:book-2-outline"/>    
                Posts
            </span>
            <span class="tab">
                <Icon icon="material-symbols:android-wifi-3-bar-rounded"/>    
                Public activity
            </span>
            <span class="tab">
                <Icon icon="material-symbols:person-2-outline-rounded"/>    
                Following
            </span>
            <span class="tab">
                <Icon icon="material-symbols:person-2-outline-rounded"/>    
                Followers
            </span>
        </div>
    </div>
</div>

<style>
    .tab-chosen {
        padding: 8px;
        color: #a4abc6;
        font-weight: 500;
        display: flex;
        flex-direction: row;
        width: fit-content;
        align-items: center;
        gap: 5px;
        font-size: 0.9rem;
        border-bottom: 2px solid #9399b2;

    }
    .tab {
        padding: 8px;
        color: #a4abc6;
        font-weight: 300;
        display: flex;
        flex-direction: row;
        width: fit-content;
        align-items: center;
        font-size: 0.9rem;
        gap: 5px;

    }
    .navbar {
        width: 800px;
        padding: 0px;
        box-sizing: border-box;
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 10px;
        border-bottom: 1px solid #38394a;
    }
    .profile_creation {
        color: #aeb9df;
        font-size: 0.9rem;
        width: 100%;
        padding: 10px;
        box-sizing: border-box;
        display: flex;
        font-weight: 300;
        flex-direction: row;
        align-items: center;
        gap: 10px;
    }

    .profile_creation :global(svg) {
        width: 18px;
        height: 18px;
        flex-shrink: 0;
    }

    .description {
        padding: 15px 10px;
        box-sizing: border-box;
        width: 100%;
        color: #aeb9df;
        font-weight: 300;
        font-size: 0.9rem;
    }
    .link {
        display: flex;
        flex-direction: row;
        gap: 10px;
        align-items: center;
    }
    .separator {
        width: 100%;
        height: 1px;
        background-color: #3f4153;
        flex-shrink: 0;
    }

    .more-btn {
        width: fit-content;
        height: fit-content;
        display: flex;
        justify-content: center;
        align-items: center;
        background-color: transparent;
        outline: none;
        border: none;
        border-radius: 5px;
        color: #c9d2f0;
        cursor: pointer;
        padding: 2px;
        transition: background-color 0.15s ease, color 0.15s ease;
    }

    .more-btn:hover {
        background-color: #29293b;
        color: #ffffff;
    }

    .more-btn :global(svg) {
        width: 24px;
        height: 24px;
    }

    .home_main {
        width: 100%;
        height: 100vh;
        background-color: #11111b;
        padding: 50px;
        box-sizing: border-box;
        display: flex;
        flex-direction: row;
        justify-content: center;
        gap: 40px;
    }

    .info_section {
        width: 310px;
        height: fit-content;
        border: 1px solid #3f4153;
        display: flex;
        flex-direction: column;
        align-items: center;
        background-color: #181825;
        border-radius: 4px;
        overflow: hidden;
    }

    .actual_content {
        width: 1000px;
    }

    .profile_picture {
        width: 270px;
        height: 270px;
        padding-top: 20px;
        box-sizing: content-box;
    }

    .profile_picture img {
        width: 100%;
        height: 100%;
        object-fit: cover;
        border-radius: 4px;
        display: block;
    }

    .basic_info {
        padding-top: 10px;
        display: flex;
        flex-direction: column;
        align-items: center;
    }

    .username {
        color: #c9d2f0;
        font-weight: 600;
        font-size: 1.1rem;
    }

    .pronouns {
        color: #8588a3;
        font-weight: 400;
        font-size: 0.9rem;
    }

    .follow_stats {
        color: #c9d2f0;
        padding-top: 15px;
        padding-bottom: 10px;
        font-weight: 400;
        font-size: 0.9rem;
        display: flex;
        flex-direction: row;
        align-items: center;
        gap: 5px;
    }

    .links {
        width: 100%;
        box-sizing: border-box;
        padding: 10px;
        display: flex;
        flex-direction: column;
        gap: 9px;
        color: #8588a3;
    }

    .links :global(svg) {
        width: 18px;
        height: 18px;
        flex-shrink: 0;
    }

    .links a {
        min-width: 0;
        color: #aeb9df;
        font-size: 0.9rem;
        font-weight: 400;
        text-decoration: none;
        white-space: nowrap;
        overflow: hidden;
        text-overflow: ellipsis;
        transition: color 0.15s ease;
    }

    .links a:hover {
        color: #ffffff;
        text-decoration: underline;
        text-underline-offset: 3px;
    }

    .links a:focus-visible {
        outline: 2px solid #7aa2f7;
        outline-offset: 3px;
        border-radius: 2px;
    }
</style>