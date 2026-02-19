import Link from "next/link";

export default function Home() {
  return (
    <div className="min-h-screen bg-gradient-to-br from-blue-50 to-indigo-100">
      <main className="container mx-auto max-w-4xl px-6 py-16">
        <div className="rounded-2xl bg-white p-8 shadow-xl md:p-12">
          <div className="mb-8">
            <h1 className="mb-4 text-5xl font-bold text-gray-900">
              Welcome to Knowledge Hub 📚
            </h1>
            <p className="text-xl text-gray-600">
              A personal knowledge management system built with modern web
              technologies
            </p>
          </div>

          <div className="mb-8 grid gap-6 md:grid-cols-2">
            <div className="rounded-xl border-2 border-blue-200 bg-gradient-to-br from-blue-50 to-blue-100 p-6">
              <h2 className="mb-3 text-2xl font-bold text-gray-900">
                🎯 Project Goals
              </h2>
              <ul className="space-y-2 text-gray-700">
                <li>• Learn multiple frontend frameworks</li>
                <li>• Practice backend API design</li>
                <li>• Build a training platform</li>
                <li>• Compare tech stacks</li>
              </ul>
            </div>

            <div className="rounded-xl border-2 border-green-200 bg-gradient-to-br from-green-50 to-green-100 p-6">
              <h2 className="mb-3 text-2xl font-bold text-gray-900">
                🛠️ Tech Stack
              </h2>
              <ul className="space-y-2 text-gray-700">
                <li>
                  • <strong>Backend:</strong> Go + PostgreSQL
                </li>
                <li>
                  • <strong>Frontend:</strong> Next.js (active)
                </li>
                <li>
                  • <strong>Monorepo:</strong> pnpm workspaces
                </li>
                <li>
                  • <strong>Styling:</strong> Tailwind CSS
                </li>
              </ul>
            </div>
          </div>

          <div className="mb-8 rounded-xl border-2 border-purple-200 bg-gradient-to-r from-purple-50 to-pink-50 p-6">
            <h2 className="mb-3 text-2xl font-bold text-gray-900">
              🚀 Quick Links
            </h2>
            <div className="flex flex-wrap gap-3">
              <Link
                href="/users-example"
                className="inline-block rounded-lg bg-blue-600 px-6 py-3 font-semibold text-white transition-colors hover:bg-blue-700"
              >
                API Example
              </Link>
              <a
                href="https://github.com/PrinceM13/knowledge-hub"
                target="_blank"
                rel="noopener noreferrer"
                className="inline-block rounded-lg bg-gray-800 px-6 py-3 font-semibold text-white transition-colors hover:bg-gray-900"
              >
                GitHub Repo
              </a>
            </div>
          </div>

          <div className="border-t-2 border-gray-200 pt-6">
            <p className="text-center text-gray-600">
              Built with ❤️ using Next.js, Go, and PostgreSQL
            </p>
          </div>
        </div>
      </main>
    </div>
  );
}
