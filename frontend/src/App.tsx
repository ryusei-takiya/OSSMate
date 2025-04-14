import { useEffect, useState } from 'react';
import './App.css'; // 相対パスで import！


type Repository = {
  full_name: string;
  description: string;
  stargazers_count: number;
  html_url: string;
};

function App() {
  const [language, setLanguage] = useState('');
  const [repos, setRepos] = useState<Repository[]>([]);
  const [page, setPage] = useState(1);

  useEffect(() => {
    const fetchRepos = async () => {
      const url = language 
      ? `/api/trending?language=${language}&page=${page}` 
      : `/api/trending?page=${page}`
      const res = await fetch(url);
      const data = await res.json();
      setRepos(data);
    };
    fetchRepos();
  }, [language, page]);

  return (
    <div className="container">
      <h1 className="title">OSSMate 🚀</h1>
      <h2 className="description" style={{ fontSize: '0.9rem' }}>
        人気のOSSリポジトリを言語ごとに検索できるツールです。
      </h2>

      <label style={{ fontWeight: 'bold', marginRight: '0.5rem' }}>
        言語で絞り込み：
        <select className="select-box" value={language} onChange={e => setLanguage(e.target.value)}>
            <option translate="no" value="">すべて</option>
            <optgroup translate="no" label="バックエンド">
              <option translate="no" value="go">Go</option>
              <option translate="no" value="rust">Rust</option>
              <option translate="no" value="python">Python</option>
              <option translate="no" value="java">Java</option>
              <option translate="no" value="c">C</option>
              <option translate="no" value="c">C++</option>
              <option translate="no" value="c#">C#</option>
              <option translate="no" value="php">PHP</option>
            </optgroup>
            <optgroup translate="no" label="フロントエンド">
              <option translate="no" value="javascript">JavaScript</option>
              <option translate="no" value="typescript">TypeScript</option>
              <option translate="no" value="html">HTML</option>
              <option translate="no" value="css">CSS</option>
            </optgroup>
          </select>
      </label>

      <ul className="repo-list">
        {repos.map(repo => (
          <li className="repo-card" key={repo.html_url}>
            <a className="repo-title" href={repo.html_url} target="_blank" rel="noreferrer">
              {repo.full_name}
            </a>
            <p className="repo-stars">
              お気に入り数：⭐ {repo.stargazers_count}
            </p>
            <p className="repo-description">
              {repo.description ?? 'No description'}
            </p>
          </li>
        ))}
      </ul>

      <div style={{ display: 'flex', gap: '1rem', marginTop: '2rem' }}>
        <button
          disabled={page === 1}
          onClick={() => setPage(prev => Math.max(prev - 1, 1))}
        >
          ◀ 前へ
        </button>
        <span>Page {page}</span>
        <button onClick={() => setPage(prev => prev + 1)}>次へ ▶</button>
      </div>

    </div>
  );
}

export default App;
