import { useSessionStore } from "@/stores/session";
import type { NavigationGuard } from "vue-router";

export function createAuthGuard(): NavigationGuard {
    return async (to, from, next) => {
        const sessionStore = useSessionStore();
        const isUserAuthenticated = await sessionStore.checkAuthentication();

        const loginPath = '/';
        const homePath = '/home';
        const noAuthPaths = ['/'];

        const isAuthPath = !noAuthPaths.includes(to.path);
        if (isAuthPath) {
            if (!isUserAuthenticated) {
                next(loginPath);
                return;
            }
        }

        const isLoginPath = to.path == loginPath;
        if(isLoginPath) {
            if(isUserAuthenticated){
                next(homePath);
                return;
            }
        }

        next();
    }
}