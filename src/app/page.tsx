'use client'

import { Edu_SA_Beginner, Roboto } from 'next/font/google'
import { sendIcon } from './icons/Icons';
import { useState } from 'react'

const edu = Edu_SA_Beginner({ subsets: ['latin'] })
const roboto = Roboto({ weight: '300', subsets: ['latin'] })

export default function Home() {
  const [longURL, setLongURL] = useState<string>('')
  const [shortUrl, setShortUrl] = useState<string>('')

  const fetchShortnedUrl = async () => {
    const request = await fetch('http://localhost:8080/shorten', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({ url: longURL })
    })

    const data = await request.json()
    console.log(data.ShortUrl)
    setShortUrl(data.ShortUrl)
  }

  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLongURL(e.target.value)
  }

  return (
    <div className="flex flex-col h-screen items-center p-2">

      <span className={`text-white text-8xl ${edu.className} justify-start p-6 mb-14`}>
        Chisai - 小さい
      </span>

      <div className='h-full flex flex-col justify-center'>
        <span className={`text-white text-6xl ${roboto.className} mb-10`}>
          Your minimalistic URL shortner.
        </span>

        <div className='flex gap-4 mb-8'>
          <input type="text" name="urlInput" id="1" placeholder=' Insert your URL here'
            className='bg-zinc-800 w-[90%] p-5 border border-gray-400 rounded-md text-white hover:border-white'
            onChange={handleInput}
          >
          </input>
          <button
            className='flex items-center justify-center bg-zinc-800 w-[10%] p-5 border border-gray-400 rounded-md text-white
            hover:border-white'
            onClick={fetchShortnedUrl}>
            {sendIcon()}
          </button>
        </div>

        <div className='flex gap-4'>
          <div className='bg-zinc-800 w-[75%] p-5 border border-gray-400 rounded-md text-white'>
            {shortUrl ? (
              <a href={shortUrl} target="_blank" rel="noopener noreferrer" className='text-blue-200 hover:underline'>
                {shortUrl}
              </a>
            ) :
              ""
            }
          </div>

          <button
            className='flex items-center justify-center bg-zinc-800 w-[25%] p-5 border border-gray-400 rounded-md text-white
            hover:border-white gap-2'>
            Copy Link {sendIcon()}
          </button>
        </div>

      </div>
    </div>
  );
}
