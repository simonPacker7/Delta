<script setup lang="ts">
import router from '@/router'
import { useSessionStore } from '@/stores/session'
import { ref } from 'vue'
import { EnvelopeIcon, LockClosedIcon } from '@heroicons/vue/24/outline'

const userEmail = ref('')
const userPassword = ref('')
const isLoading = ref(false)
const errorMessage = ref('')
const sendLogin = useSessionStore().loginUser

const login = async () => {
    isLoading.value = true
    errorMessage.value = ''
    try {
        await sendLogin({ email: userEmail.value, password: userPassword.value })
        router.push('/home')
    } catch (error: any) {
        errorMessage.value = error?.response?.data?.error || 'Invalid email or password'
    } finally {
        isLoading.value = false
    }
}
</script>

<template>
    <form @submit.prevent="login" class="auth-form">
        <div class="form-group">
            <label for="loginEmail" class="form-label">Email</label>
            <div class="input-wrapper">
                <EnvelopeIcon class="input-icon" />
                <input
                    v-model="userEmail"
                    type="email"
                    class="form-input"
                    id="loginEmail"
                    placeholder="you@example.com"
                    required
                    :disabled="isLoading"
                />
            </div>
        </div>

        <div class="form-group">
            <label for="loginPassword" class="form-label">Password</label>
            <div class="input-wrapper">
                <LockClosedIcon class="input-icon" />
                <input
                    v-model="userPassword"
                    type="password"
                    class="form-input"
                    id="loginPassword"
                    placeholder="••••••••"
                    required
                    :disabled="isLoading"
                />
            </div>
        </div>

        <div v-if="errorMessage" class="error-message">
            {{ errorMessage }}
        </div>

        <button type="submit" class="submit-btn" :disabled="isLoading">
            <span v-if="isLoading" class="loading-spinner"></span>
            <span v-else>Sign In</span>
        </button>
    </form>
</template>

<style scoped>
.auth-form {
    display: flex;
    flex-direction: column;
    gap: 1.25rem;
}

.form-group {
    display: flex;
    flex-direction: column;
    gap: 0.5rem;
}

.form-label {
    font-size: 0.875rem;
    font-weight: 500;
    color: #d1d5db;
    font-family: monospace;
}

.input-wrapper {
    position: relative;
    display: flex;
    align-items: center;
}

.input-icon {
    position: absolute;
    left: 1rem;
    width: 1.25rem;
    height: 1.25rem;
    color: #6b7280;
    pointer-events: none;
}

.form-input {
    width: 100%;
    padding: 0.875rem 1rem 0.875rem 3rem;
    background: rgba(255, 255, 255, 0.05);
    border: 1px solid rgba(255, 255, 255, 0.1);
    border-radius: 0.75rem;
    color: #ffffff;
    font-size: 1rem;
    font-family: inherit;
    transition: all 0.2s ease;
}

.form-input::placeholder {
    color: #6b7280;
}

.form-input:focus {
    outline: none;
    border-color: #6366f1;
    box-shadow: 0 0 0 3px rgba(99, 102, 241, 0.2);
    background: rgba(255, 255, 255, 0.08);
}

.form-input:disabled {
    opacity: 0.6;
    cursor: not-allowed;
}

.error-message {
    background: rgba(239, 68, 68, 0.1);
    border: 1px solid rgba(239, 68, 68, 0.3);
    color: #fca5a5;
    padding: 0.75rem 1rem;
    border-radius: 0.5rem;
    font-size: 0.875rem;
    text-align: center;
}

.submit-btn {
    width: 100%;
    padding: 0.875rem 1.5rem;
    background: linear-gradient(135deg, #6366f1 0%, #8b5cf6 100%);
    border: none;
    border-radius: 0.75rem;
    color: #ffffff;
    font-size: 1rem;
    font-weight: 600;
    font-family: monospace;
    cursor: pointer;
    transition: all 0.2s ease;
    display: flex;
    align-items: center;
    justify-content: center;
    gap: 0.5rem;
    margin-top: 0.5rem;
}

.submit-btn:hover:not(:disabled) {
    transform: translateY(-2px);
    box-shadow: 0 10px 20px rgba(99, 102, 241, 0.3);
}

.submit-btn:active:not(:disabled) {
    transform: translateY(0);
}

.submit-btn:disabled {
    opacity: 0.7;
    cursor: not-allowed;
}

.loading-spinner {
    width: 1.25rem;
    height: 1.25rem;
    border: 2px solid rgba(255, 255, 255, 0.3);
    border-top-color: #ffffff;
    border-radius: 50%;
    animation: spin 0.8s linear infinite;
}

@keyframes spin {
    to {
        transform: rotate(360deg);
    }
}

@media (max-width: 640px) {
    .form-input {
        padding: 0.75rem 0.875rem 0.75rem 2.75rem;
        font-size: 0.9rem;
    }

    .input-icon {
        left: 0.875rem;
        width: 1.125rem;
        height: 1.125rem;
    }

    .submit-btn {
        padding: 0.75rem 1.25rem;
    }
}
</style>