import axios from "axios";

const api = axios.create({
  baseURL: import.meta.env.VITE_API_URL || "http://localhost:8080/api",
  timeout: 10000,
  headers: {
    "Content-Type": "application/json",
  },
});

api.interceptors.request.use(
  (config) => {
    const token = localStorage.getItem("authToken");
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
    return config;
  },
  (error) => Promise.reject(error)
);

api.interceptors.response.use(
  (response) => response.data,
  (error) => {
    console.error("API Error:", error);
    return Promise.reject(error);
  }
);

export const authAPI = {
  login: async (credentials: { username: string; password: string }) => {
    return api.post("/auth/login", credentials);
  },

  register: async (userData: {
    username: string;
    email: string;
    phone: string;
    password: string;
  }) => {
    return api.post("/auth/register", userData);
  },

  logout: async () => {
    return api.post("/auth/logout");
  },

  resetPassword: async (email: string) => {
    return api.post("/auth/reset-password", { email });
  },
};

export const userAPI = {
  getProfile: async () => {
    return api.get("/user/profile");
  },

  updateProfile: async (profileData: Record<string, unknown>) => {
    return api.put("/user/profile", profileData);
  },

  getSubscription: async () => {
    return api.get("/user/subscription");
  },

  updateSubscription: async (planId: string) => {
    return api.put("/user/subscription", { planId });
  },
};

export const reviewAPI = {
  createReview: async (data: {
    topic: string;
    impactFactorRange?: { min: number; max: number };
    yearRange?: { startYear: number; endYear: number };
  }) => {
    return api.post("/reviews", data);
  },

  getReview: async (id: string) => {
    return api.get(`/reviews/${id}`);
  },

  updateReview: async (id: string, data: Record<string, unknown>) => {
    return api.put(`/reviews/${id}`, data);
  },

  deleteReview: async (id: string) => {
    return api.delete(`/reviews/${id}`);
  },

  getUserReviews: async (page = 1, limit = 10) => {
    return api.get(`/reviews?page=${page}&limit=${limit}`);
  },

  generateReviewContent: async (reviewId: string) => {
    return api.post(`/reviews/${reviewId}/generate`);
  },

  checkGenerationStatus: async (taskId: string) => {
    return api.get(`/reviews/generation/${taskId}`);
  },
};

export const writingAPI = {
  startWriting: async (payload: {
    title: string;
    topic: string;
    locale?: string;
    pubmedQuery?: string;
    includeFilters?: boolean;
  }) => {
    return api.post("/writing", payload);
  },
};

export default api;
