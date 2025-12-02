<script setup lang="ts">

</script>

<template>
    <header>
        <h1 class="title">Delta</h1>
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
    background-color: var(--vt-c-black-soft);
    font-family: monospace;
}

.title {
    font-weight: 500;
    font-size: 3.5rem;
    color: beige;
}
</style>