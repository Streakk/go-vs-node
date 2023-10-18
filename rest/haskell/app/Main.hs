{-# LANGUAGE DeriveGeneric #-}
{-# LANGUAGE DeriveAnyClass #-}
{-# LANGUAGE OverloadedStrings #-}

module Main where

import Data.Aeson (ToJSON, FromJSON)
import Data.Text (Text)
import GHC.Generics (Generic)
import Web.Scotty

data Numbers = Numbers
  { values :: ![Double]
  } deriving (Show, Eq, Generic, FromJSON)

data Result = Result
  { totalSum    :: Double
  , average     :: Double
  , totalProd   :: Double
  } deriving (Show, Eq, Generic, ToJSON)

compute :: [Double] -> Result
compute vals = Result s avg prod
  where
    s = sum vals
    avg = s / fromIntegral (length vals)
    prod = product vals

main :: IO ()
main = scotty 8080 $ do
  post "/compute" $ do
    nums <- jsonData :: ActionM Numbers
    let res = compute (values nums)
    json res
