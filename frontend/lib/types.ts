export interface User {
  id: number;
  email: string;
  firstname?: string;
  lastname?: string;
}

export interface Event {
  id: number;
  name: string;
  description: string;
  location: string;
  datetime: string;
  user_id: number;
}

export interface LoginCredentials {
  email: string;
  password: string;
}

export interface SignupData {
  email: string;
  password: string;
  firstname?: string;
  lastname?: string;
}

export interface ApiResponse<T> {
  data?: T;
  message?: string;
  error?: string;
}
