<script setup lang="ts">
import { computed, type Component } from 'vue'
import { useRouter, useRoute } from 'vue-router'

const props = defineProps({
    label: { type: String, required: true },
    routePath: { type: String, required: false },
    icon: { type: Object as () => Component, required: false },
    click: { type: Function, required: false },
})

const router = useRouter()
const route = useRoute()

const isActiveRoute = computed(() => {
    return props.routePath && route.path.startsWith(props.routePath)
})

const navigate = () => {
    if (props.click) {
        props.click()
    } else if (props.routePath) {
        router.push({ path: props.routePath })
    }
}
</script>

<template>
    <button 
        class="nav-item" 
        :class="{ 'active': isActiveRoute }" 
        @click="navigate"
    >
        <component v-if="icon" :is="icon" class="nav-icon" />
        <span class="nav-label">{{ label }}</span>
    </button>
</template>

<style scoped>
.nav-item {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.625rem 1rem;
    background: transparent;
    border: none;
    border-radius: 0.5rem;
    color: #9ca3af;
    font-family: monospace;
    font-size: 0.9rem;
    font-weight: 500;
    cursor: pointer;
    transition: all 0.2s ease;
    white-space: nowrap;
}

.nav-item:hover {
    background: rgba(255, 255, 255, 0.08);
    color: #ffffff;
}

.nav-item.active {
    background: rgba(99, 102, 241, 0.15);
    color: #a5b4fc;
}

.nav-item.active .nav-icon {
    color: #6366f1;
}

.nav-icon {
    width: 1.25rem;
    height: 1.25rem;
    flex-shrink: 0;
    transition: color 0.2s ease;
}

.nav-label {
    line-height: 1;
}

/* Mobile styles - stretch full width */
@media (max-width: 900px) {
    .nav-item {
        width: 100%;
        padding: 0.875rem 1rem;
        border-radius: 0.75rem;
    }
}
</style>