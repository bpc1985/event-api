import Link from 'next/link';
import { Button } from '@/components/ui/button';

export default function HomePage() {
  return (
    <div className="container mx-auto px-4 py-16">
      <div className="max-w-3xl mx-auto text-center space-y-8">
        <h1 className="text-5xl font-bold tracking-tight">
          Welcome to EventApp
        </h1>
        <p className="text-xl text-muted-foreground">
          Discover, create, and manage events with ease. Join a community of event enthusiasts
          and never miss out on exciting opportunities.
        </p>
        <div className="flex gap-4 justify-center">
          <Link href="/events">
            <Button size="lg">Browse Events</Button>
          </Link>
          <Link href="/signup">
            <Button size="lg" variant="outline">
              Get Started
            </Button>
          </Link>
        </div>
      </div>
    </div>
  );
}
