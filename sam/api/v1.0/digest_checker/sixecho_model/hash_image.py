from sqlalchemy import Column, DateTime, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class HashImage(Base):
    __tablename__ = 'hash_images'
    id = Column(String, primary_key=True)
    hash_type = Column(String, primary_key=True)
    digital_content_id = Column(String)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
