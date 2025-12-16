# Event Management Frontend

A modern Next.js frontend for the Event Management API built with React 19, TypeScript, and TailwindCSS.

## Features

- Browse and view events
- User authentication (signup/login)
- Event registration
- View event attendees
- Responsive design with shadcn/ui components

## Getting Started

### Prerequisites

- Node.js 18+ (installed via nvm)
- The Event Management API running on `http://localhost:8080`

### Installation

```bash
# Install dependencies
npm install

# Run the development server
npm run dev
```

Open [http://localhost:3000](http://localhost:3000) in your browser.

### Build for Production

```bash
npm run build
npm run start
```

## Environment Variables

Create a `.env.local` file in the root directory:

```env
NEXT_PUBLIC_API_URL=http://localhost:8080
JWT_SECRET=your-secret-key-here
```

## Project Structure

```
frontend/
├── app/                  # Next.js App Router pages
│   ├── events/          # Events pages
│   ├── login/           # Login page
│   ├── signup/          # Signup page
│   └── page.tsx         # Home page
├── components/           # React components
│   ├── auth/            # Authentication components
│   ├── events/          # Event components
│   └── ui/              # Reusable UI components (shadcn/ui)
├── lib/                  # Utilities and API client
│   ├── api.ts           # API client with axios
│   ├── auth.ts          # Authentication utilities
│   ├── types.ts         # TypeScript types
│   └── utils.ts         # Helper functions
└── public/              # Static assets
```

## Tech Stack

- **Framework**: Next.js 15 with App Router
- **Language**: TypeScript
- **Styling**: TailwindCSS
- **UI Components**: shadcn/ui (built on Radix UI)
- **HTTP Client**: axios
- **Authentication**: JWT with jose library
- **Icons**: Lucide React

## API Integration

The frontend connects to the Event Management API running on `localhost:8080`. Make sure the API is running before starting the frontend.

### Available Endpoints

- `GET /events` - List all events
- `GET /events/:id` - Get event details
- `GET /events/:id/attendees` - Get event attendees
- `POST /signup` - Register new user
- `POST /login` - Authenticate user
- `POST /events` - Create event (authenticated)
- `PUT /events/:id` - Update event (authenticated)
- `DELETE /events/:id` - Delete event (authenticated)
- `POST /events/:id/register` - Register for event (authenticated)
- `DELETE /events/:id/register` - Unregister from event (authenticated)
