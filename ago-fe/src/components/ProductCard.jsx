import React from 'react';
import { Pencil, Trash2 } from 'lucide-react';
import { useNavigate } from 'react-router-dom';

export function ProductCard({ id, name, description, onDelete }) {
  const navigate = useNavigate();

  const handleEdit = () => {
    navigate(`/edit/${id}`);
  };

  return (
    <div className="bg-white rounded-lg shadow-md p-6 flex flex-col">
      <h3 className="text-xl font-semibold mb-2">{name}</h3>
      <p className="text-gray-600 flex-grow mb-4">{description}</p>
      <div className="flex justify-end space-x-2">
        <button
          onClick={handleEdit}
          className="p-2 text-blue-600 hover:text-blue-800 transition-colors"
        >
          <Pencil size={20} />
        </button>
        <button
          onClick={onDelete}
          className="p-2 text-red-600 hover:text-red-800 transition-colors"
        >
          <Trash2 size={20} />
        </button>
      </div>
    </div>
  );
}
