"use client";

import { Input } from '@/components/ui/form/Input';
import { Button } from '@/components/ui/button/Button';
import { useState } from 'react';
import { searchFormData } from '@/types/search';
import { fetchMapInfo } from '@/services/searchService';

export default function Home() {
  const [form, setForm] = useState<searchFormData>({ address: "", l_limit_travel_expenses: 0, u_limit_travel_expenses: 0 });
  const [message, setMessage] = useState("");

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setForm({ ...form, [e.target.name]: e.target.value });
  };

  const handleSubmit = async () => {
    try {
      const res = await fetchMapInfo(form);
      setMessage(`APIコール完了, res:${JSON.stringify(res)}`);
    } catch (err) {
      console.error(err);
      setMessage(`APIコールが失敗しました:${err}`);
    }
  };

  return (
    <div>
      <div>
        住所：
        <Input
          type='text'
          name='address'
          placeholder='住所を入力'
          onChange={handleChange}
        />
      </div>
      <div>
        金額：
        <Input
          type='number'
          name='l_limit_travel_expenses'
          placeholder='下限値'
          className='w-20'
          onChange={handleChange}
        />
        ~
        <Input
          type='number'
          name='u_limit_travel_expenses'
          placeholder='上限値'
          className='w-20'
          onChange={handleChange}
        />
      </div>
      <div>
        <Button onClick={handleSubmit} className='active:bg-blue-700'>
          検索
        </Button>
      </div>
      {message && <p className="mt-3 text-green-700">{message}</p>}
    </div>
  );
}
