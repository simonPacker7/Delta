import axios from "axios";
import { defineStore } from "pinia";
import { ref } from "vue";

export const useSessionStore = defineStore("session", () => {
    const email = ref('');
    const displayName = ref('');
    const id = ref('');
    const authenticated = ref(false);
    const fresh = ref(true);

    const $reset = () => {
        email.value = "";
        displayName.value = "";
        id.value = "";
        authenticated.value = false;
        fresh.value = true;
    }

    const syncUserProfile = async () => {
        try {
            const res = await axios.get('/api/user/profile')
            if (!res || res.status != 200 || !res.data || !res.data.email || !res.data.name) {
                authenticated.value = false;
                return
            }
            email.value = res.data.email
            displayName.value = res.data.name
            id.value = res.data.id
            authenticated.value = true
        } catch (err) {
            authenticated.value = false;
            console.log(err)
        }
    }

    const checkAuthentication = async () => {
        if (fresh.value) {
            await syncUserProfile()
            fresh.value = false
        }
        return authenticated.value
    }

    const registerUser = async (input: { email: string, name: string, password: string }) => {
        await axios.post('/api/auth/register', input);
        await syncUserProfile();
    }

    const loginUser = async (input: { email: string, password: string }) => {
        await axios.post('/api/auth/login', input);
        await syncUserProfile();
    }

    const logoutUser = async () => {
        await axios.post('/api/auth/logout');
        $reset();
    }

    return { email, displayName, id, authenticated, checkAuthentication, registerUser, loginUser, logoutUser }
})