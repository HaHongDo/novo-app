import { useContext, useState } from 'react';
import Link from 'next/link';
import { MangaContext } from '../../Context/MangaContext';
import NULL_CONSTANTS from '../../utilities/nullConstants';
import DisplayImg from '../../components/displayImg';
import WEB_CONSTANTS from '../../utilities/constants';
import ReactPaginate from 'react-paginate';
export async function getServerSideProps(context) {
  const server = WEB_CONSTANTS.SERVER;
  const { params } = context;
  const { id } = params;
  const response = await fetch(`${server}/author/${id}`);
  const data = await response.json();

  return {
    props: {
      author: data,
    },
  };
}

export default function Author({ author }) {
  const { server } = useContext(MangaContext);
  const [pageNumber, setPageNumber] = useState(0);
  const bookPerPage = 4;
  const pageVisited = pageNumber * bookPerPage;
  const displayDatas =
    author.books.length > 0 &&
    author.books
      .slice(pageVisited, pageVisited + bookPerPage)
      .map((book) => (
        <Link
          href={`/mangas/${book.id}`}
          passHref
          key={book.id}
        >
          <div
            className="col-lg-3 col-12"
            data-aos="fade-up"
          >
            <DisplayImg
              srcImg={
                book.image
                  ? `${server}/image/${book.image}`
                  : NULL_CONSTANTS.BOOK_GROUP_IMAGE
              }
              text={'Chap ' + book.latestChapter}
              title={book.title}
              bgColor="green"
            ></DisplayImg>
          </div>
        </Link>
      ));
  const pageCount = Math.ceil(
    author.books.length / bookPerPage
  );
  const changePage = ({ selected }) => {
    setPageNumber(selected);
  };

  return (
    <div>
      <div
        className="author-gradient container-fluid"
        data-aos="fade-in"
      >
        <img
          className="rounded-circle"
          data-aos="fade-down"
          src={
            author.avatar
              ? `${server}/image/${author.avatar}`
              : NULL_CONSTANTS.BOOK_GROUP_IMAGE
          }
        ></img>
      </div>
      <div className="mt-3 container">
        <div className="row">
          <div className="col-lg-2 col-12 author-detail"></div>
          <div className="col-lg-9 col-12">
            <h3>{author.name}</h3>
            <br />
            <h5>Mô tả</h5>
            <p style={{ wordBreak: 'break-all' }}>
              {author.description}
            </p>
            <h5>Sáng tác</h5>
            <div className="row">{displayDatas}</div>
            <div className="d-flex justify-content-center">
              <ReactPaginate
                breakLabel="..."
                previousLabel="Trước"
                nextLabel="Sau"
                pageCount={pageCount}
                onPageChange={changePage}
                pageClassName="page-item"
                pageLinkClassName="page-link"
                previousClassName="page-item"
                previousLinkClassName="page-link"
                nextClassName="page-item"
                nextLinkClassName="page-link"
                breakClassName="page-item"
                breakLinkClassName="page-link"
                containerClassName="pagination"
                activeClassName="active"
                renderOnZeroPageCount={null}
              ></ReactPaginate>
            </div>
          </div>
        </div>
      </div>
    </div>
  );
}
