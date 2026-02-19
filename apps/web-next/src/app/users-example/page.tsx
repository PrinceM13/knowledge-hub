export default async function UsersExample() {
  // Example: Fetch users from API
  // Uncomment when your Go API is running
  /*
  import { api } from "@/lib/api";
  
  const users = await api.users.listUsers({ limit: 10 });
  
  return (
    <div>
      <h1>Users</h1>
      <ul>
        {users.map((user) => (
          <li key={user.email}>{user.name} - {user.email}</li>
        ))}
      </ul>
    </div>
  );
  */

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

        <div className="mb-6 rounded-lg border-2 border-blue-300 bg-blue-100 p-6">
          <h2 className="mb-3 text-xl font-bold text-gray-900">
            To use the API:
          </h2>
          <ol className="list-inside list-decimal space-y-2 text-base text-gray-800">
            <li>Start your Go API server</li>
            <li>Uncomment the code above</li>
            <li>Refresh this page</li>
          </ol>
        </div>

        <div className="rounded-lg border-2 border-gray-300 bg-gray-100 p-6">
          <h2 className="mb-3 text-xl font-bold text-gray-900">
            Type-safe API calls:
          </h2>
          <pre className="overflow-x-auto rounded border border-gray-300 bg-white p-4 text-sm">
            <code className="text-gray-900">{`import { api } from '@/lib/api';

// List users
const users = await api.users.listUsers({ limit: 10 });

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
