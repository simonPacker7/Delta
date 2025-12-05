<script setup lang="ts">
import router from '@/router';
import { useSessionStore } from '@/stores/session';
import { ref } from 'vue';

const userEmail = ref('');
const userPassword = ref('');
const sendLogin = useSessionStore().loginUser;

const login = async() => {
    await sendLogin({ email: userEmail.value, password: userPassword.value });
    router.push('/home');
}
</script>

<template>
    <form @submit.prevent="login">
        <div class="title-group">
            <label class="form-title mono">Login</label>
        </div>
        <div class="form-group">
            <label for="loginEmail" class="form-label mono">Email</label>
            <input v-model="userEmail" type="email" class="form-control" id="loginEmail" required />
        </div>
        <div class="form-group">
            <label for="loginPassword" class="form-label mono">Password</label>
            <input v-model="userPassword" type="password" class="form-control" id="loginPassword" required />
        </div>
        <button type="submit" class="btn mono">Sign In</button>
    </form>
</template>

<style scoped>
form>button {
    width: 100%;
    margin-top: 27px;
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