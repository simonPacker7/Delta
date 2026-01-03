<script setup lang="ts">
import { ref } from 'vue'
import navBar from "./NavBar.vue"
import userNavBar from "./UserNavBar.vue"
import { useSessionStore } from "@/stores/session"
import { Bars3Icon, XMarkIcon } from '@heroicons/vue/24/outline'

const mobileMenuOpen = ref(false)

const toggleMobileMenu = () => {
    mobileMenuOpen.value = !mobileMenuOpen.value
}

const closeMobileMenu = () => {
    mobileMenuOpen.value = false
}
</script>

<template>
    <header>
        <div class="header-content">
            <!-- Logo -->
            <div class="logo-section">
                <h1 class="title">Delta</h1>
            </div>

            <!-- Desktop Navigation -->
            <navBar class="desktop-nav" />
            <userNavBar class="desktop-user-nav" />

            <!-- Mobile Menu Button -->
            <button 
                class="mobile-menu-btn"
                @click="toggleMobileMenu"
                :aria-expanded="mobileMenuOpen"
                aria-label="Toggle navigation menu"
            >
                <Bars3Icon v-if="!mobileMenuOpen" class="menu-icon" />
                <XMarkIcon v-else class="menu-icon" />
            </button>
        </div>

        <!-- Mobile Navigation Overlay -->
        <Transition name="slide">
            <div 
                v-if="mobileMenuOpen" 
                class="mobile-nav-overlay"
                @click="closeMobileMenu"
            >
                <div class="mobile-nav-content" @click.stop>
                    <navBar class="mobile-nav" @navigate="closeMobileMenu" />
                    <div class="mobile-divider"></div>
                    <userNavBar class="mobile-user-nav" @navigate="closeMobileMenu" />
                </div>
            </div>
        </Transition>
    </header>
</template>

<style scoped>
header {
    position: fixed;
    top: 0;
    left: 0;
    right: 0;
    z-index: 1000;
    background: rgba(30, 30, 35, 0.85);
    backdrop-filter: blur(20px);
    border-bottom: 1px solid rgba(255, 255, 255, 0.08);
}

.header-content {
    display: flex;
    align-items: center;
    justify-content: space-between;
    padding: 0 2rem;
    height: 70px;
    max-width: 1400px;
    margin: 0 auto;
}

.logo-section {
    display: flex;
    align-items: center;
}

.title {
    font-family: monospace;
    font-weight: 600;
    font-size: 1.75rem;
    margin: 0;
    background: linear-gradient(135deg, #ffffff 0%, #a5b4fc 50%, #6366f1 100%);
    -webkit-background-clip: text;
    -webkit-text-fill-color: transparent;
    background-clip: text;
}

/* Desktop Navigation */
.desktop-nav {
    display: flex;
    margin-left: 3rem;
}

.desktop-user-nav {
    display: flex;
    margin-left: auto;
}

/* Mobile Menu Button */
.mobile-menu-btn {
    display: none;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.5rem;
    padding: 0.5rem;
    cursor: pointer;
    transition: all 0.2s ease;
}

.mobile-menu-btn:hover {
    background: rgba(255, 255, 255, 0.1);
}

.menu-icon {
    width: 1.5rem;
    height: 1.5rem;
    color: #ffffff;
}

/* Mobile Navigation Overlay */
.mobile-nav-overlay {
    display: none;
    position: fixed;
    top: 70px;
    left: 0;
    right: 0;
    bottom: 0;
    background: rgba(0, 0, 0, 0.6);
    backdrop-filter: blur(4px);
    z-index: 999;
}

.mobile-nav-content {
    background: rgba(30, 30, 35, 0.98);
    border-bottom: 1px solid rgba(255, 255, 255, 0.1);
    padding: 1rem;
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.mobile-divider {
    height: 1px;
    background: rgba(255, 255, 255, 0.1);
    margin: 0.5rem 0;
}

/* Slide transition */
.slide-enter-active,
.slide-leave-active {
    transition: opacity 0.2s ease;
}

.slide-enter-active .mobile-nav-content,
.slide-leave-active .mobile-nav-content {
    transition: transform 0.2s ease;
}

.slide-enter-from,
.slide-leave-to {
    opacity: 0;
}

.slide-enter-from .mobile-nav-content,
.slide-leave-to .mobile-nav-content {
    transform: translateY(-10px);
}

/* Responsive */
@media (max-width: 900px) {
    .desktop-nav,
    .desktop-user-nav {
        display: none;
    }

    .mobile-menu-btn {
        display: flex;
    }

    .mobile-nav-overlay {
        display: block;
    }
}

@media (max-width: 480px) {
    .header-content {
        padding: 0 1rem;
        height: 60px;
    }

    .title {
        font-size: 1.5rem;
    }

    .mobile-nav-overlay {
        top: 60px;
    }
}
</style>