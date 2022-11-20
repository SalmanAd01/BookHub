import { Stack } from '@mui/material'
import React from 'react'
import BooksCard from './BooksCard'
import Hero from './Hero'

const Home = () => {
  return (
    <Stack>
    <Hero></Hero>
    <BooksCard></BooksCard>
    </Stack>
  )
}

export default Home