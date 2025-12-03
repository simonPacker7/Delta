<script setup lang="ts">
import HeaderView from "@/components/header/Header.vue"
import { RouterView } from 'vue-router'
</script>

<template>
  <HeaderView />
  <main>
    <RouterView />
  </main>
</template>

<style scoped>
main {
  /* Adjust for header height (TODO: Adjust to actual height not this random amount) */
  margin-top: 130px;
  width: 100%;
}
</style>

<style>
body {
  background-color: #282828;
  color: white;
}
.mono {
  font-family: monospace !important;
}
</style>