<script setup lang="ts">
import navItem from "./NavItem.vue"
import { HomeIcon, PlayIcon, ChartBarIcon, TrophyIcon } from '@heroicons/vue/24/outline'

const emit = defineEmits(['navigate'])

const onNavigate = () => {
    emit('navigate')
}
</script>

<template>
    <nav class="nav-bar">
        <navItem routePath="/home" label="Home" :icon="HomeIcon" @click="onNavigate" />
        <navItem routePath="/play" label="Play" :icon="PlayIcon" @click="onNavigate" />
        <navItem routePath="/statistics" label="Statistics" :icon="ChartBarIcon" @click="onNavigate" />
        <navItem routePath="/leaderboard" label="Leaderboard" :icon="TrophyIcon" @click="onNavigate" />
    </nav>
</template>

<style scoped>
.nav-bar {
    display: flex;
    align-items: center;
    gap: 0.5rem;
}

/* Mobile styles applied via parent class */
:deep(.mobile-nav) .nav-bar {
    flex-direction: column;
    align-items: stretch;
}
</style>