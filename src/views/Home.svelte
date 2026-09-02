<script lang="ts">
    import { api } from '../lib/api';

    async function getAvatar() {
        const response = await api('/media/avatar.jpg');

        if (!response.ok) {
            throw new Error('Nie udało się pobrać avatara');
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
    </div>
    <div class="actual_content">

    </div>
</div>
<style>
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
        width: 320px;
        border: 1px solid #3f4153;
        display: flex;
        flex-direction: column;
        align-items: center;
    }
    .actual_content {
        width: 1000px;
    }
    .profile_picture {
        width: 280px;
        height: 280px;
        padding-top: 20px;
    }
    .profile_picture img {
        width: 100%;
        border-radius: 5px;
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
        color: #c9d2f0;
        font-weight: 400;
        font-size: 0.9rem;
    }
</style>
