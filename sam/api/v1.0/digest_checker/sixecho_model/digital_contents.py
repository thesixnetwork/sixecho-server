from sqlalchemy import Column, DateTime, Integer, String, Text
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class DigitalContent(Base):
    __tablename__ = 'digital_contents'
    id = Column(Integer, primary_key=True)
    api_key_id = Column(String)
    category_id = Column(Integer)
    publisher_id = Column(Integer)
    digest = Column(Text)
    sha256 = Column(String)
    size_file = Column(String)
    title = Column(String)
    author = Column(String)
    country_of_origin = Column(String)
    language = Column(String)
    publish_date = Column(DateTime)
    paperback = Column(Integer)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
