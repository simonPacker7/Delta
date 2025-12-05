<script setup lang="ts">
import router from '@/router';
import { useSessionStore } from '@/stores/session';
import { ref } from 'vue';

const userEmail = ref('');
const userDisplayName = ref('');
const userPassword = ref('');
const sendRegister = useSessionStore().registerUser;

const register = async () => {
    await sendRegister({ email: userEmail.value, name: userDisplayName.value, password: userPassword.value })
    router.push('/home')
}
</script>

<template>
    <form @submit.prevent="register">
        <div class="title-group">
            <label class="form-title mono">Register</label>
        </div>
        <div class="form-group">
            <label for="registerEmail" class="form-label mono">Email</label>
            <input v-model="userEmail" type="email" class="form-control" id="registerEmail"
                required />
        </div>
        <div class="form-group">
            <label for="registerDisplayName" class="form-label mono">Display name</label>
            <input v-model="userDisplayName" class="form-control" type="text" id="registerDisplayName" required
                maxlength="16" />
        </div>
        <div class="form-group">
            <label for="registerPassword" class="form-label mono">Password</label>
            <input v-model="userPassword" type="password" class="form-control" id="registerPassword" required
                minlength="8" />
        </div>
        <button type="submit" class="btn mono">Join</button>
    </form>
</template>

<style scoped>
form>button {
    width: 100%
}
.title-group{
    display: flex;
    justify-content: center;
    align-items: center;
    margin-bottom: 10px;
}
.form-group {
    margin-bottom: 7px;
}
.form-label {
    margin-right: 7px;
}
.form-title {
    font-weight: bold;
}
</style>