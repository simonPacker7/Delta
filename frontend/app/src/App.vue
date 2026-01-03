<script setup lang="ts">
import HeaderView from "@/components/header/Header.vue"
import { RouterView, useRoute } from 'vue-router'
import { computed } from 'vue'

const route = useRoute()
const isLandingPage = computed(() => route.path === '/')
</script>

<template>
  <HeaderView v-if="!isLandingPage" />
  <main :class="{ 'with-header': !isLandingPage }">
    <RouterView />
  </main>
</template>

<style scoped>
main {
  width: 100%;
  min-height: 100vh;
}

main.with-header {
  /* Adjust for header height (70px desktop, 60px mobile) */
  padding-top: 70px;
}

@media (max-width: 480px) {
  main.with-header {
    padding-top: 60px;
  }
}
</style>

<style>
*, *::before, *::after {
  box-sizing: border-box;
}

html, body {
  margin: 0;
  padding: 0;
  overflow-x: hidden;
}

body {
  background-color: #282828;
  color: white;
}
.mono {
  font-family: monospace !important;
}
</style>