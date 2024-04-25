import { useState } from "react";
import NewsCard from "../components/NewsCard";

function Home() {
  const newsItems = [
    {
      ID: 0,
      CreatedAt: "0001-01-01T00:00:00Z",
      UpdatedAt: "0001-01-01T00:00:00Z",
      DeletedAt: null,
      Title: "Stock market today: Live updates - CNBC",
      Description:
        "Stocks ended Thursday near the flat line after a fresh round of inflation data reflected an uptick in consumer prices for December.",
      Content:
        "Stocks ended Thursday near the flat line after a fresh round of inflation data reflected an uptick in consumer prices for December.\r\nThe Nasdaq Composite closed at the flat line, settling at 14,970.1\u2026 [+2669 chars]",
      URL: "https://www.cnbc.com/2024/01/10/stock-market-today-live-updates.html",
      ImageURL:
        "https://image.cnbcfm.com/api/v1/image/107281049-1691038968344-gettyimages-1588215007-0j6a6715_atzuolmb.jpeg?v=1704928252&w=1920&h=1080",
      PublishedAt: "2024-01-11T21:15:00Z",
      Source: "CNBC",
      Category: "Business",
      Keywords: [
        "CNBC Stocks ended",
        "Stocks ended Thursday",
        "Stock market today",
        "inflation data reflected",
        "Live updates",
        "CNBC Stocks",
        "prices for December",
        "ended Thursday",
        "market today",
        "Stock market",
        "Stocks ended",
        "flat line",
        "fresh round",
        "round of inflation",
        "inflation data",
        "data reflected",
        "reflected an uptick",
        "uptick in consumer",
        "consumer prices",
        "Live",
      ],
      Likes: null,
      Dislikes: null,
    },
    {
      ID: 0,
      CreatedAt: "0001-01-01T00:00:00Z",
      UpdatedAt: "0001-01-01T00:00:00Z",
      DeletedAt: null,
      Title:
        "Ecuador\u2019s chaotic turn to drug and narco violence, explained - Vox.com",
      Description:
        "Ecuador was known for peace, but it has become one of the most violent countries in South America.",
      Content:
        "Ecuador, according to its president Daniel Noboa, is now in a state of war. Earlier this week he had announced a state of emergency after the leader of one of the countrys top two gangs escaped from \u2026 [+10457 chars]",
      URL: "https://www.vox.com/world-politics/2024/1/11/24034891/ecuador-drugs-cocaine-cartels-violence-murder-daniel-naboa-columbia-crime",
      ImageURL:
        "https://cdn.vox-cdn.com/thumbor/svlre_mpFzHYZrvji7F13fnEysU=/0x588:8192x4877/fit-in/1200x630/cdn.vox-cdn.com/uploads/chorus_asset/file/25219855/GettyImages_1915423205.jpg",
      PublishedAt: "2024-01-11T21:45:00Z",
      Source: "Vox",
      Category: "World",
      Keywords: [
        "South America",
        "Vox.com Ecuador",
        "countries in South",
        "narco violence",
        "chaotic turn",
        "turn to drug",
        "drug and narco",
        "violent countries",
        "explained",
        "Vox.com",
        "America",
        "Ecuador \u2019s chaotic",
        "Ecuador",
        "South",
        "violence",
        "peace",
        "chaotic",
        "turn",
        "drug",
        "narco",
      ],
      Likes: null,
      Dislikes: null,
    },
  ];
  const [currentSlide, setCurrentSlide] = useState(0);

  const nextSlide = () => {
    setCurrentSlide((prev) => (prev + 1) % newsItems.length);
  };

  const prevSlide = () => {
    setCurrentSlide((prev) => (prev === 0 ? newsItems.length - 1 : prev - 1));
  };

  return (
    <div>
      <h1>Top Headlines</h1>
      <div className="carousel">
        <button onClick={prevSlide} className="carousel-control left">
          &lt;
        </button>
        <div className="carousel-slide">
          <NewsCard news={newsItems[currentSlide]} />
        </div>
        <button onClick={nextSlide} className="carousel-control right">
          &gt;
        </button>
      </div>
    </div>
  );
}

export default Home;
