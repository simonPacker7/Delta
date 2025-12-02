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
  /* Adjust for header height */
  margin-top: 76px;
  width: 100%;
}
</style>
