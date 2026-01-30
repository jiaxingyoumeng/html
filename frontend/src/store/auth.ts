import { defineStore } from "pinia";

import { authAPI } from "../services/apiService";

export interface User {
  id: string;
  username: string;
  email: string;
  avatar?: string;
  subscription?: {
    level: string;
    expiryDate: string;
    totalReviews: number;
    usedReviews: number;
  };
}

interface AuthState {
  isAuthenticated: boolean;
  user: User | null;
  loading: boolean;
  showLoginModal: boolean;
  showRegisterModal: boolean;
}

const loadUserFromStorage = (): User | null => {
  const storedUser = localStorage.getItem("user");
  return storedUser ? (JSON.parse(storedUser) as User) : null;
};

export const useAuthStore = defineStore("auth", {
  state: (): AuthState => ({
    isAuthenticated: localStorage.getItem("isAuthenticated") === "true",
    user: loadUserFromStorage(),
    loading: false,
    showLoginModal: false,
    showRegisterModal: false,
  }),

  actions: {
    async login(credentials: { username: string; password: string }) {
      this.loading = true;
      try {
        const response = await authAPI.login(credentials);
        this.isAuthenticated = true;
        this.user = response.user as User;

        localStorage.setItem("isAuthenticated", "true");
        localStorage.setItem("user", JSON.stringify(response.user));
        localStorage.setItem("authToken", response.token);

        return true;
      } catch (error) {
        console.error("Login failed:", error);
        return false;
      } finally {
        this.loading = false;
      }
    },

    logout() {
      this.isAuthenticated = false;
      this.user = null;
      localStorage.removeItem("isAuthenticated");
      localStorage.removeItem("user");
      localStorage.removeItem("authToken");
    },

    async register(userData: {
      username: string;
      email: string;
      phone: string;
      password: string;
    }) {
      this.loading = true;
      try {
        const response = await authAPI.register(userData);
        this.isAuthenticated = true;
        this.user = response.user as User;

        localStorage.setItem("isAuthenticated", "true");
        localStorage.setItem("user", JSON.stringify(response.user));
        localStorage.setItem("authToken", response.token);

        return true;
      } catch (error) {
        console.error("Register failed:", error);
        return false;
      } finally {
        this.loading = false;
      }
    },

    setShowLoginModal(show: boolean) {
      this.showLoginModal = show;
    },

    setShowRegisterModal(show: boolean) {
      this.showRegisterModal = show;
    },
  },
});
