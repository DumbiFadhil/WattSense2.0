import { useState } from "react";
import axios from "axios";

const AIChatConsultantPage = () => {
  const [query, setQuery] = useState("");
  const [response, setResponse] = useState("");
  const [isLoading, setIsLoading] = useState(false);
  const [chatHistory, setChatHistory] = useState([]);

  const handleChat = async () => {
    if (!query.trim()) {
      alert("Please enter a query.");
      return;
    }

    const newChatHistory = [...chatHistory, { type: 'user', message: query }];
    setChatHistory(newChatHistory);
    setIsLoading(true);

    try {
      const res = await axios.post("http://localhost:8080/chat", { query });
      const aiResponse = res.data.answer;
      
      setResponse(aiResponse);
      setChatHistory(prev => [...prev, { type: 'ai', message: aiResponse }]);
    } catch (error) {
      console.error("Error querying chat:", error);
      setResponse("An error occurred while processing your request.");
    } finally {
      setIsLoading(false);
      setQuery("");
    }
  };

  return (
    <div className="container mx-auto mt-10 max-w-2xl">
      <h1 className="text-3xl font-bold mb-6 text-center">AI Consultant</h1>
      <div className="bg-white shadow-md rounded-lg p-6">
        <div className="mb-4 h-96 overflow-y-auto border rounded p-4">
          {chatHistory.map((msg, index) => (
            <div 
              key={index} 
              className={`mb-2 p-2 rounded ${msg.type === 'user' ? 'bg-blue-100 text-right' : 'bg-green-100 text-left'}`}
            >
              {msg.message}
            </div>
          ))}
        </div>
        <div className="flex">
          <input
            type="text"
            value={query}
            onChange={(e) => setQuery(e.target.value)}
            placeholder="Ask your AI consultant..."
            className="flex-grow p-2 border rounded-l"
            onKeyPress={(e) => e.key === 'Enter' && handleChat()}
          />
          <button 
            onClick={handleChat} 
            disabled={isLoading}
            className="bg-green-500 text-white p-2 rounded-r hover:bg-green-600 transition duration-300 disabled:opacity-50"
          >
            {isLoading ? 'Thinking...' : 'Send'}
          </button>
        </div>
      </div>
    </div>
  );
};

export default AIChatConsultantPage;
