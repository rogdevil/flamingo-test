import React from 'react';
import { BrowserRouter, Routes, Route } from 'react-router-dom';
import { ProductList } from './pages/ProductList';
import { ProductForm } from './pages/ProductForm';

function App() {
  return (
    <BrowserRouter>
      <div className="min-h-screen bg-gray-50">
        <Routes>
          <Route path="/" element={<ProductList />} />
          <Route path="/create" element={<ProductForm />} />
          <Route path="/edit/:id" element={<ProductForm />} />
        </Routes>
      </div>
    </BrowserRouter>
  );
}

export default App;
