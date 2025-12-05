<script setup lang="ts">
import { computed } from 'vue'
import { useRouter, useRoute } from 'vue-router'
const props = defineProps({
    label: { type: String, required: true },
    routePath: { type: String, required: false },
    icon: { type: String, required: false },
    click: { type: Function, required: false },
})

const router = useRouter()
const route = useRoute()

const isActiveRoute = computed(() => {
    return props.routePath && route.path.startsWith(props.routePath)
})

const navigate = () => {
    router.push({ path: props.routePath })
}
</script>

<template>
    <button class="nav-bar-item" :class="{ 'active-route': isActiveRoute }" :onclick="click ? click : navigate">
        {{ label }}
    </button>
</template>

<style scoped>
.nav-bar-item {
    font-weight: 300;
    font-size: 1.25rem;
    color: black;
    margin-left: 3rem;
}

.nav-bar-item:hover {
    border-color: black;
    color: black;
    border-radius: 0.375rem;
}

.active-route {
    border-bottom-color: #0d6efd;
    border-radius: 1px;
}

.icon {
    color: #5c8bd1;
}
</style>