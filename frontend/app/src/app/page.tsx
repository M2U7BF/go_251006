import { Input } from '@/components/ui/form/Input';
import { Button } from '@/components/ui/button/Button';

export default function Home() {
  return (
    <div>
      <div>
        住所：
        <Input
          type='text'
          placeholder='住所を入力'
        />
      </div>
      <div>
        金額：
        <Input
          type='number'
          placeholder='下限値'
          className='w-20'
        />
        ~
        <Input
          type='number'
          placeholder='上限値'
          className='w-20'
        />
      </div>
      <div>
        <Button>
          検索
        </Button>
      </div>
    </div>
  );
}
