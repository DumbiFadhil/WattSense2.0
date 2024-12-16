import { Link } from "react-router-dom";

const Navigation = () => (
  <nav className="bg-blue-600 p-4 text-white">
    <div className="container mx-auto flex justify-between items-center">
      <h1 className="text-2xl font-bold">Data Analysis Toolkit</h1>
      <div className="space-x-4">
        <Link to="/" className="hover:bg-blue-700 px-3 py-2 rounded">Home</Link>
        <Link to="/data-analyzer" className="hover:bg-blue-700 px-3 py-2 rounded">Data Analyzer</Link>
        <Link to="/ai-consultant" className="hover:bg-blue-700 px-3 py-2 rounded">AI Consultant</Link>
      </div>
    </div>
  </nav>
);

export default Navigation;
