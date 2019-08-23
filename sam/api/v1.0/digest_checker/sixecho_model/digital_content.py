from sqlalchemy import Column, DateTime, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class DigitalContent(Base):
    __tablename__ = 'digital_contents'
    id = Column(String, primary_key=True)
    api_key_id = Column(String)
    category_id = Column(Integer)
    publisher_id = Column(Integer)
    digital_content_id = Column(String)
    title = Column(String)
    digest = Column(String)
    sha256 = Column(String)
    size_file = Column(String)
    content_type = Column(String)
    meta_media = Column(String)
    author = Column(String)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
