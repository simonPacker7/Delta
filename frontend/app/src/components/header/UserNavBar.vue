<script setup lang="ts">
import { useRouter } from 'vue-router'
import { useSessionStore } from "@/stores/session"
import { UserCircleIcon, ArrowRightOnRectangleIcon } from '@heroicons/vue/24/outline'
import navItem from "./NavItem.vue"

const emit = defineEmits(['navigate'])

const sessionStore = useSessionStore()
const router = useRouter()

const logout = async () => {
    await sessionStore.logoutUser()
    emit('navigate')
    await router.replace("/")
}

const onNavigate = () => {
    emit('navigate')
}
</script>

<template>
    <nav class="user-nav">
        <div class="user-info">
            <UserCircleIcon class="user-avatar" />
            <span class="user-name">{{ sessionStore.displayName }}</span>
        </div>
        <div class="user-actions">
            <navItem routePath="/profile" label="Profile" @click="onNavigate" />
            <button class="logout-btn" @click="logout">
                <ArrowRightOnRectangleIcon class="logout-icon" />
                <span>Sign out</span>
            </button>
        </div>
    </nav>
</template>

<style scoped>
.user-nav {
    display: flex;
    align-items: center;
    gap: 1rem;
}

.user-info {
    display: flex;
    align-items: center;
    gap: 0.5rem;
    padding: 0.375rem 0.75rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 2rem;
}

.user-avatar {
    width: 1.5rem;
    height: 1.5rem;
    color: #6366f1;
}

.user-name {
    font-family: monospace;
    font-size: 0.875rem;
    font-weight: 500;
    color: #e5e7eb;
    max-width: 120px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
}

.user-actions {
    display: flex;
    align-items: center;
    gap: 0.25rem;
}

.logout-btn {
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
}

.logout-btn:hover {
    background: rgba(239, 68, 68, 0.1);
    color: #fca5a5;
}

.logout-btn:hover .logout-icon {
    color: #ef4444;
}

.logout-icon {
    width: 1.25rem;
    height: 1.25rem;
    transition: color 0.2s ease;
}

/* Mobile styles */
@media (max-width: 900px) {
    .user-nav {
        flex-direction: column;
        align-items: stretch;
        gap: 0.5rem;
    }

    .user-info {
        justify-content: center;
        padding: 0.75rem 1rem;
        border-radius: 0.75rem;
    }

    .user-name {
        max-width: none;
    }

    .user-actions {
        flex-direction: column;
    }

    .logout-btn {
        width: 100%;
        padding: 0.875rem 1rem;
        border-radius: 0.75rem;
        justify-content: flex-start;
    }
}
</style>