<script setup lang="ts">
import navBar from "./NavBar.vue"
import userNavBar from "./UserNavBar.vue"
import { useSessionStore } from "@/stores/session"

const sessionStore = useSessionStore()
</script>

<template>
    <header>
        <h1 class="title">δ Delta</h1>
        <navBar v-if="sessionStore.authenticated" />
        <userNavBar v-if="sessionStore.authenticated" />
    </header>
</template>

<style scoped>
header {
    position: fixed;
    top: 0px;
    left: 0px;
    right: 0px;
    z-index: 1000;
    display: flex !important;
    align-items: center;
    justify-content: flex-start;
    flex-direction: row;
    padding-left: 5%;
    /* background-color: rgb(54, 124, 147); */
    background-color: #2c3e50;
    font-family: monospace;
}

.title {
    font-weight: 500;
    font-size: 3.5rem;
    color: beige;
}
</style>