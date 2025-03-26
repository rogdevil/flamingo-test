import React, { useEffect, useState } from 'react';
import { Plus } from 'lucide-react';
import { Link } from 'react-router-dom';
import { ProductCard } from '../components/ProductCard';

export function ProductList() {
  const [products, setProducts] = useState([]);

  useEffect(() => {
    fetchProducts();
  }, []);

  const fetchProducts = async () => {
    const requestOptions = {
      method: 'GET',
      redirect: 'follow',
      headers: {
        Accept: 'application/json',
        'Content-Type': 'application/json',
      },
      mode: 'cors',
      credentials: 'include',
    };

    let apiError;

    const data = await fetch(
      'http://localhost:3322/products/paginated?page=1&pageSize=5',
      requestOptions,
    )
      .then((response) => response.json())
      .then((result) => result)
      .catch((error) => (apiError = error));

    if (apiError) {
      console.error('Error fetching products:', apiError);
      return;
    }

    console.log(data, 'data is');

    setProducts(data || []);
  };

  const handleDelete = async (id) => {
    console.log('delete');

    setProducts(products.filter((product) => product.id !== id));
  };

  return (
    <div className="container mx-auto px-4 py-8">
      <div className="flex justify-between items-center mb-8">
        <h1 className="text-3xl font-bold text-gray-900">Products</h1>
        <Link
          to="/create"
          className="bg-blue-600 text-white px-4 py-2 rounded-lg flex items-center gap-2 hover:bg-blue-700 transition-colors"
        >
          <Plus size={20} />
          Add Product
        </Link>
      </div>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        {products.map((product) => (
          <ProductCard key={product.id} {...product} onDelete={() => handleDelete(product.id)} />
        ))}
      </div>
    </div>
  );
}
