'use client'

import { Edu_SA_Beginner, Roboto } from 'next/font/google'
import { sendIcon, copyIcon } from './icons/Icons';
import { useState } from 'react'

const edu = Edu_SA_Beginner({ subsets: ['latin'] })
const roboto = Roboto({ weight: '300', subsets: ['latin'] })

export default function Home() {
  const [longURL, setLongURL] = useState<string>('')
  const [shortUrl, setShortUrl] = useState<string>('')
  const [error, setError] = useState<String>('')

  const fetchShortnedUrl = async () => {
    try {
      const response = await fetch('http://localhost:8080/shorten', {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json'
        },
        body: JSON.stringify({ url: longURL })
      })

      if (!response.ok) {
        const errorText = await response.text()
        throw new Error(errorText)
      }

      const data = await response.json()
      setShortUrl(data.ShortUrl)
      setError('')
      console.log(shortUrl)
    }

    catch (error: any) {
      console.log("There was a issue fetching data from the server: ")
      setError(error.message)
      setShortUrl('')
    }
  }

  const handleInput = (e: React.ChangeEvent<HTMLInputElement>) => {
    setLongURL(e.target.value)
  }

  const enterKeySubmit = (e: React.KeyboardEvent<HTMLInputElement>) => {
    if (e.key === 'Enter') {
      fetchShortnedUrl()
    }
  }

  return (
    <div className="flex flex-col h-screen items-center p-2">
      <span className={`text-white text-8xl ${edu.className} mt-8`}>
        Chisai 小さい
      </span>

      <div className='h-full flex flex-col mt-44'>
        <span className={`text-white text-6xl ${roboto.className} mb-12`}>
          Your minimalistic URL shortner.
        </span>

        <div className='flex gap-4 mb-8'>
          <input type="text" name="urlInput" id="1" placeholder=' Insert your URL here'
            className='bg-zinc-800 w-[85%] p-5 border border-gray-400 rounded-md text-white hover:border-white'
            onChange={handleInput}
            onKeyDown={enterKeySubmit}
            autoComplete='off'
          >
          </input>
          <button
            className='flex items-center justify-center bg-zinc-800 w-[15%] p-5 border border-gray-400 rounded-md text-white
            hover:border-white'
            onClick={fetchShortnedUrl}>
            {sendIcon()}
          </button>
        </div>

        <div className='flex gap-4 mb-8'>
          <div className='bg-zinc-800 w-[75%] p-5 border border-gray-400 rounded-md text-white hover:border-white'>
            {shortUrl ? (
              <a href={shortUrl} target="_blank" rel="noopener noreferrer" className='text-blue-200 hover:underline'>
                {shortUrl}
              </a>
            ) :
              ""
            }
          </div>

          <button
            className='flex justify-center items-center bg-zinc-800 w-[25%] p-5 border border-gray-400 rounded-md text-white
            hover:border-white gap-2'
            onClick={() => { navigator.clipboard.writeText(shortUrl) }}
          >
            <span className='hidden sm:flex'>
              Copy Link
            </span>
            {copyIcon()}
          </button>
        </div>

        {error ? (
          <div className="bg-red-600 text-white p-4 rounded-md">
            {error}
          </div>
        ) : ''}

      </div>
    </div>
  );
}
