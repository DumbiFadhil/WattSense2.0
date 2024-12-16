import { BrowserRouter as Router, Routes, Route, Navigate } from 'react-router-dom';
import Navigation from './components/Navigation';
import HomePage from './pages/HomePage';
import DataAnalyzerPage from './pages/DataAnalyzerPage';
import AIChatConsultantPage from './pages/AIChatConsultantPage';

function App() {
  return (
    <Router>
      <div className="min-h-screen bg-gray-100">
        <Navigation />
        <Routes>
          <Route path="/" element={<HomePage />} />
          <Route path="/data-analyzer" element={<DataAnalyzerPage />} />
          <Route path="/ai-consultant" element={<AIChatConsultantPage />} />
          <Route path="*" element={<Navigate to="/" replace />} />
        </Routes>
      </div>
    </Router>
  );
}

export default App;
