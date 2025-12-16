'use client';

import { useEffect, useState } from 'react';
import { useParams } from 'next/navigation';
import { eventsApi } from '@/lib/api';
import { Event, User } from '@/lib/types';
import { formatDate } from '@/lib/utils';
import { Button } from '@/components/ui/button';
import { Card, CardContent, CardDescription, CardHeader, CardTitle } from '@/components/ui/card';
import { Calendar, MapPin, Users } from 'lucide-react';
import { getToken } from '@/lib/auth';

export default function EventDetailPage() {
  const params = useParams();
  const eventId = Number(params.id);

  const [event, setEvent] = useState<Event | null>(null);
  const [attendees, setAttendees] = useState<User[]>([]);
  const [loading, setLoading] = useState(true);
  const [error, setError] = useState('');
  const [registering, setRegistering] = useState(false);
  const isAuthenticated = !!getToken();

  useEffect(() => {
    const fetchEventData = async () => {
      try {
        const [eventResponse, attendeesResponse] = await Promise.all([
          eventsApi.getById(eventId),
          eventsApi.getAttendees(eventId),
        ]);
        setEvent(eventResponse.data);
        setAttendees(attendeesResponse.data);
      } catch (err: any) {
        setError(err.response?.data?.message || 'Failed to load event');
      } finally {
        setLoading(false);
      }
    };

    if (eventId) {
      fetchEventData();
    }
  }, [eventId]);

  const handleRegister = async () => {
    if (!isAuthenticated) {
      window.location.href = '/login';
      return;
    }

    setRegistering(true);
    try {
      await eventsApi.register(eventId);
      const attendeesResponse = await eventsApi.getAttendees(eventId);
      setAttendees(attendeesResponse.data);
    } catch (err: any) {
      alert(err.response?.data?.message || 'Failed to register');
    } finally {
      setRegistering(false);
    }
  };

  if (loading) {
    return (
      <div className="container mx-auto px-4 py-16">
        <div className="text-center">Loading event...</div>
      </div>
    );
  }

  if (error || !event) {
    return (
      <div className="container mx-auto px-4 py-16">
        <div className="text-center text-red-500">{error || 'Event not found'}</div>
      </div>
    );
  }

  return (
    <div className="container mx-auto px-4 py-16">
      <div className="max-w-4xl mx-auto space-y-8">
        <Card>
          <CardHeader>
            <CardTitle className="text-3xl">{event.name}</CardTitle>
            <CardDescription>{event.description}</CardDescription>
          </CardHeader>
          <CardContent className="space-y-6">
            <div className="flex items-center text-muted-foreground">
              <Calendar className="mr-2 h-5 w-5" />
              <span>{formatDate(event.datetime)}</span>
            </div>
            <div className="flex items-center text-muted-foreground">
              <MapPin className="mr-2 h-5 w-5" />
              <span>{event.location}</span>
            </div>
            <div className="flex items-center text-muted-foreground">
              <Users className="mr-2 h-5 w-5" />
              <span>{attendees.length} attendees</span>
            </div>
            <Button
              onClick={handleRegister}
              disabled={registering}
              className="w-full md:w-auto"
            >
              {registering ? 'Registering...' : 'Register for Event'}
            </Button>
          </CardContent>
        </Card>

        <Card>
          <CardHeader>
            <CardTitle>Attendees ({attendees.length})</CardTitle>
          </CardHeader>
          <CardContent>
            {attendees.length === 0 ? (
              <p className="text-muted-foreground">No attendees yet. Be the first to register!</p>
            ) : (
              <ul className="space-y-2">
                {attendees.map((attendee) => (
                  <li key={attendee.id} className="flex items-center gap-2">
                    <div className="w-8 h-8 rounded-full bg-primary/10 flex items-center justify-center">
                      <span className="text-sm font-medium">
                        {attendee.firstname?.[0] || attendee.email[0].toUpperCase()}
                      </span>
                    </div>
                    <span>
                      {attendee.firstname && attendee.lastname
                        ? `${attendee.firstname} ${attendee.lastname}`
                        : attendee.email}
                    </span>
                  </li>
                ))}
              </ul>
            )}
          </CardContent>
        </Card>
      </div>
    </div>
  );
}
