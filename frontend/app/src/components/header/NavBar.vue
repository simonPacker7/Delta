<script setup lang="ts">
import navItem from "./NavItem.vue"
</script>

<template>
    <nav>
        <navItem routePath="/home" label="Home" icon="fa-house" />
        <navItem routePath="/play" label="Play" icon="fa-gamepad" />
        <navItem routePath="/statistics" label="My Statistics" icon="fa-chart-line" />
        <navItem routePath="/leaderboard" label="Leaderboard" icon="fa-chart-bar" />
    </nav>
</template>