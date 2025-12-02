<script setup lang="ts">
import { ref } from 'vue';

const userEmail = ref('');
const userPassword = ref('');
</script>

<template>
    <form>
        <div class="mb-3">
            <label for="loginEmail" class="form-label mono">Email</label>
            <input v-model="userEmail" type="email" class="form-control" id="loginEmail" required />
        </div>
        <div class="mb-3">
            <label for="loginPassword" class="form-label mono">Password</label>
            <input v-model="userPassword" type="password" class="form-control" id="loginPassword" required />
        </div>
        <button type="submit" class="btn btn-outline-dark">Sign In</button>
    </form>
</template>

<style scoped>
form>button {
    width: 100%
}
</style>