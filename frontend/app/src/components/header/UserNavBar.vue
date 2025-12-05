<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useSessionStore } from "@/stores/session"
import navItem from "./NavItem.vue"

const sessionStore = useSessionStore()
const router = useRouter()

const logout = async () => {
    await sessionStore.logoutUser()
    await router.replace("/")
}

</script>

<template>
    <nav class="usernav">
        <navItem routePath="/profile" :label="sessionStore.displayName" icon="fa-user" />
        <navItem label="Sign out" icon="fa-arrow-right-from-bracket" :click="logout" />
    </nav>
</template>

<style scoped>
.usernav {
    margin-left: auto;
    margin-right: 30px;
}

.usernav>button {
    margin-left: 10px;
}
</style>