<script setup lang="ts">
import { ref } from 'vue';

const userEmail = ref('');
const userDisplayName = ref('');
const userPassword = ref('');

</script>

<template>
    <form>
        <div class="mb-3">
            <label for="registerEmail" class="form-label mono">Email</label>
            <input v-model="userEmail" type="email" class="form-control" placeholder="me@company.com" id="registerEmail"
                required />
        </div>
        <div class="mb-3">
            <label for="registerDisplayName" class="form-label mono">Display name</label>
            <input v-model="userDisplayName" class="form-control" type="text" id="registerDisplayName" required
                maxlength="16" />
        </div>
        <div class="mb-3">
            <label for="registerPassword" class="form-label mono">Password</label>
            <input v-model="userPassword" type="password" class="form-control" id="registerPassword" required
                minlength="8" />
        </div>
        <button type="submit" class="btn btn-outline-dark">Join</button>
    </form>
</template>

<style scoped>
form>button {
    width: 100%
}
</style>