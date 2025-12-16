# Frontend Setup Complete

The Next.js frontend has been successfully created and configured!

## What Was Created

### Project Structure
```
frontend/
├── app/                          # Next.js App Router pages
│   ├── events/                  # Events listing and detail pages
│   │   ├── [id]/page.tsx       # Dynamic event detail page
│   │   └── page.tsx             # Events listing page
│   ├── login/page.tsx           # Login page
│   ├── signup/page.tsx          # Signup page
│   ├── layout.tsx               # Root layout with navigation
│   ├── page.tsx                 # Home page
│   └── globals.css              # Global styles with shadcn/ui theme
├── components/
│   ├── auth/                    # Authentication components
│   │   ├── login-form.tsx      # Login form component
│   │   └── signup-form.tsx     # Signup form component
│   ├── events/
│   │   └── event-card.tsx      # Event card component
│   └── ui/                      # shadcn/ui components
│       ├── button.tsx
│       ├── card.tsx
│       ├── input.tsx
│       ├── label.tsx
│       └── textarea.tsx
├── lib/                          # Utilities and API client
│   ├── api.ts                   # Axios API client with all endpoints
│   ├── auth.ts                  # JWT authentication utilities
│   ├── types.ts                 # TypeScript type definitions
│   └── utils.ts                 # Helper functions (formatDate, cn)
└── Configuration files
    ├── package.json             # Dependencies and scripts
    ├── tsconfig.json            # TypeScript configuration
    ├── tailwind.config.ts       # Tailwind CSS configuration
    ├── next.config.ts           # Next.js configuration
    └── .env.local               # Environment variables

```

## Quick Start

### 1. Start the Backend API
First, make sure your Go API is running:
```bash
# In the root directory (event-api/)
go run .
```
The API should be running on http://localhost:8080

### 2. Start the Frontend Dev Server
In a new terminal:
```bash
cd frontend
npm run dev
```

The frontend will be available at **http://localhost:3000**

### 3. Test the Application
1. Open http://localhost:3000 in your browser
2. Click "Sign Up" to create a new account
3. Login with your credentials
4. Browse events at http://localhost:3000/events
5. Click on an event to see details and register

## Available Pages

- **/** - Home page with welcome message
- **/events** - Browse all available events
- **/events/[id]** - View event details and attendees
- **/login** - User login
- **/signup** - User registration

## Features Implemented

### Authentication
- JWT token-based authentication
- Token stored in localStorage
- Login and signup forms with validation
- Password confirmation on signup

### Events
- Display all events in a responsive grid
- View event details with attendees list
- Register for events (requires authentication)
- See attendee count and list

### UI/UX
- Modern, responsive design
- shadcn/ui components (built on Radix UI)
- Tailwind CSS styling
- Loading states and error handling
- Accessible components

## API Integration

The frontend connects to these backend endpoints:

### Public Endpoints
- `GET /events` - List all events
- `GET /events/:id` - Get event details
- `GET /events/:id/attendees` - Get event attendees
- `POST /signup` - Register new user
- `POST /login` - Authenticate user

### Protected Endpoints (require JWT token)
- `POST /events` - Create event
- `PUT /events/:id` - Update event
- `DELETE /events/:id` - Delete event
- `POST /events/:id/register` - Register for event
- `DELETE /events/:id/register` - Unregister from event

## Environment Variables

The `.env.local` file contains:
```env
NEXT_PUBLIC_API_URL=http://localhost:8080
JWT_SECRET=your-secret-key-here
```

**Note**: Make sure the `JWT_SECRET` in your Go backend matches the one used here.

## Build Commands

```bash
# Development server
npm run dev

# Production build
npm run build

# Start production server
npm run start

# Lint code
npm run lint
```

## Tech Stack

- **Next.js 15** - React framework with App Router
- **React 19** - UI library
- **TypeScript** - Type safety
- **Tailwind CSS** - Utility-first CSS
- **shadcn/ui** - Reusable components (Radix UI + Tailwind)
- **Axios** - HTTP client
- **jose** - JWT verification and signing
- **Lucide React** - Icon library

## Next Steps

You can now:
1. Customize the UI theme in `app/globals.css` and `tailwind.config.ts`
2. Add more pages (e.g., user dashboard, create event form)
3. Implement event creation UI
4. Add event update/delete functionality for event owners
5. Add user profile page
6. Implement event search and filtering

## Troubleshooting

### Port Already in Use
If port 3000 is already in use:
```bash
# Specify a different port
PORT=3001 npm run dev
```

### API Connection Issues
- Ensure the Go API is running on port 8080
- Check that CORS is properly configured in the Go backend
- Verify the `NEXT_PUBLIC_API_URL` in `.env.local`

### Build Errors
If you encounter build errors:
```bash
# Clean install
rm -rf node_modules package-lock.json
npm install
```

## Success! 🎉

Your frontend is now ready to use with your Event Management API!
