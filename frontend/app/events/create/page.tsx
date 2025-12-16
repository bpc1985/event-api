'use client';

import { useEffect } from 'react';
import { useRouter } from 'next/navigation';
import { getToken } from '@/lib/auth';
import { CreateEventForm } from '@/components/events/create-event-form';

export default function CreateEventPage() {
  const router = useRouter();

  useEffect(() => {
    const token = getToken();
    if (!token) {
      router.push('/login');
    }
  }, [router]);

  return (
    <div className="container mx-auto px-4 py-16 flex items-center justify-center">
      <CreateEventForm />
    </div>
  );
}
