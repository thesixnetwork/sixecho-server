from sqlalchemy import Column, DateTime, Integer, String
from sqlalchemy.ext.declarative import declarative_base

Base = declarative_base()


class Publisher(Base):
    __tablename__ = 'publishers'
    id = Column(Integer, primary_key=True)
    first_name = Column(String)
    last_name = Column(String)
    created_at = Column(DateTime)
    updated_at = Column(DateTime)
