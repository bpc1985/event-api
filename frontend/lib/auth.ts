import { jwtVerify, SignJWT, type JWTPayload as JoseJWTPayload } from 'jose';

const JWT_SECRET = new TextEncoder().encode(
  process.env.JWT_SECRET || 'your-secret-key-here'
);

export interface JWTPayload {
  userId: number;
  email: string;
  exp?: number;
}

export async function verifyToken(token: string): Promise<JWTPayload | null> {
  try {
    const { payload } = await jwtVerify(token, JWT_SECRET);
    // Check if payload has required custom fields
    if (payload && typeof payload === 'object' && 'userId' in payload && 'email' in payload) {
      return payload as unknown as JWTPayload;
    }
    return null;
  } catch (error) {
    console.error('Token verification failed:', error);
    return null;
  }
}

export async function generateToken(userId: number, email: string): Promise<string> {
  const token = await new SignJWT({ userId, email })
    .setProtectedHeader({ alg: 'HS256' })
    .setExpirationTime('24h')
    .sign(JWT_SECRET);
  return token;
}

export function getToken(): string | null {
  if (typeof window === 'undefined') return null;
  return localStorage.getItem('token');
}

export function setToken(token: string): void {
  if (typeof window === 'undefined') return;
  localStorage.setItem('token', token);
}

export function removeToken(): void {
  if (typeof window === 'undefined') return;
  localStorage.removeItem('token');
}

export async function isAuthenticated(): Promise<boolean> {
  const token = getToken();
  if (!token) return false;

  const payload = await verifyToken(token);
  return payload !== null;
}
