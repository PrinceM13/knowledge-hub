import type { UserListItem } from "@repo/api-contract/types";
import { api } from "@/lib/api";

export default async function UsersExample() {
  // Fetch real users from Go API
  let users: UserListItem[] = [];
  let error: string | null = null;

  try {
    users = await api.users.listUsers({ limit: 10 });
  } catch (err) {
    error = err instanceof Error ? err.message : "Failed to fetch users";
    console.error("Failed to fetch users:", err);
  }

  return (
    <div className="min-h-screen bg-white p-8">
      <div className="mx-auto max-w-4xl">
        <h1 className="mb-6 text-3xl font-bold text-gray-900">
          API Contract Example
        </h1>
        <p className="mb-6 text-lg text-gray-700">
          This page demonstrates how to use the{" "}
          <code className="rounded bg-gray-200 px-2 py-1 font-mono text-sm text-gray-900">
            @repo/api-contract
          </code>{" "}
          package.
        </p>

        {/* Users List Example */}
        <div className="mb-6 rounded-lg border-2 border-green-300 bg-green-50 p-6">
          <h2 className="mb-4 text-2xl font-bold text-gray-900">
            📋 Users List {error && "⚠️"}
          </h2>
          <p className="mb-3 text-sm text-gray-600">
            {error
              ? `Error: ${error}`
              : `Showing ${users.length} user(s) from the Go API`}
          </p>
          {users.length > 0 ? (
            <div className="space-y-2">
              {users.map((user) => (
                <div
                  key={user.email}
                  className="rounded border border-green-200 bg-white p-3"
                >
                  <div className="font-semibold text-gray-900">{user.name}</div>
                  <div className="text-sm text-gray-600">{user.email}</div>
                </div>
              ))}
            </div>
          ) : (
            !error && (
              <div className="rounded bg-yellow-50 p-4 text-gray-700">
                No users found. Try creating one using the API!
              </div>
            )
          )}
        </div>

        <div className="mb-6 rounded-lg border-2 border-blue-300 bg-blue-100 p-6">
          <h2 className="mb-3 text-xl font-bold text-gray-900">
            ✅ Connected to Real API!
          </h2>
          <p className="mb-2 text-base text-gray-800">
            This page is now fetching data from your Go API at{" "}
            <code className="rounded bg-white px-1 text-sm">
              http://localhost:8080/api/v1
            </code>
          </p>
          <p className="text-sm text-gray-700">
            The data you see above is coming from your PostgreSQL database! 🎉
          </p>
        </div>

        <div className="rounded-lg border-2 border-gray-300 bg-gray-100 p-6">
          <h2 className="mb-3 text-xl font-bold text-gray-900">
            💻 Type-safe API calls:
          </h2>
          <pre className="overflow-x-auto rounded border border-gray-300 bg-white p-4 text-sm">
            <code className="text-gray-900">{`import { api } from '@/lib/api';
import type { UserListItem } from '@repo/api-contract/types';

// List users (with type safety!)
const users: UserListItem[] = await api.users.listUsers({ 
  limit: 10,
  offset: 0 
});

// Create user
const newUser = await api.users.createUser({
  email: 'test@example.com',
  name: 'Test User',
});

// Get user by ID
const user = await api.users.getUserById(1);`}</code>
          </pre>
        </div>
      </div>
    </div>
  );
}
