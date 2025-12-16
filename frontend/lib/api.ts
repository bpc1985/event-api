import axios from 'axios';
import type { Event, User, LoginCredentials, SignupData } from './types';

const API_BASE_URL = process.env.NEXT_PUBLIC_API_URL || 'http://localhost:8080';

const apiClient = axios.create({
  baseURL: API_BASE_URL,
  headers: {
    'Content-Type': 'application/json',
  },
});

// Add request interceptor to include auth token
apiClient.interceptors.request.use((config) => {
  if (typeof window !== 'undefined') {
    const token = localStorage.getItem('token');
    if (token) {
      config.headers.Authorization = `Bearer ${token}`;
    }
  }
  return config;
});

// Events API
export const eventsApi = {
  getAll: () => apiClient.get<Event[]>('/events'),
  getById: (id: number) => apiClient.get<Event>(`/events/${id}`),
  create: (event: Omit<Event, 'id' | 'user_id'>) => apiClient.post<{ message: string; event: Event }>('/events', event),
  update: (id: number, event: Partial<Event>) => apiClient.put<{ message: string }>(`/events/${id}`, event),
  delete: (id: number) => apiClient.delete<{ message: string }>(`/events/${id}`),
  register: (id: number) => apiClient.post<{ message: string }>(`/events/${id}/register`),
  unregister: (id: number) => apiClient.delete<{ message: string }>(`/events/${id}/register`),
  getAttendees: (id: number) => apiClient.get<User[]>(`/events/${id}/attendees`),
};

// Users API
export const usersApi = {
  signup: (data: SignupData) => apiClient.post<{ message: string }>('/signup', data),
  login: (credentials: LoginCredentials) => apiClient.post<{ message: string; token: string }>('/login', credentials),
  getEventsByAttendee: (id: number) => apiClient.get<Event[]>(`/attendees/${id}/events`),
};

export default apiClient;
