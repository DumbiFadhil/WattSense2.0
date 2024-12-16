import { Link } from 'react-router-dom';

const HomePage = () => {
  return (
    <div className="container mx-auto mt-10 text-center">
      <h1 className="text-4xl font-bold mb-6">Welcome to Data Analysis Toolkit</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 gap-6 max-w-2xl mx-auto">
        <Link 
          to="/data-analyzer" 
          className="bg-blue-500 text-white p-6 rounded-lg shadow-md hover:bg-blue-600 transition duration-300"
        >
          <h2 className="text-2xl font-semibold mb-4">Data Analyzer</h2>
          <p>Upload and analyze your datasets with advanced tools</p>
        </Link>
        <Link 
          to="/ai-consultant" 
          className="bg-green-500 text-white p-6 rounded-lg shadow-md hover:bg-green-600 transition duration-300"
        >
          <h2 className="text-2xl font-semibold mb-4">AI Consultant</h2>
          <p>Get AI-powered insights and recommendations</p>
        </Link>
      </div>
    </div>
  );
};

export default HomePage;
